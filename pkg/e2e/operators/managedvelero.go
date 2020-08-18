package operators

import (
	"context"
	"fmt"
	"time"

	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/osde2e/pkg/common/helper"
	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
)

var _ = ginkgo.Describe("[Suite: operators] [OSD] Managed Velero Operator", func() {
	var operatorName = "managed-velero-operator"
	var operatorNamespace string = "openshift-velero"
	var operatorLockFile string = "managed-velero-operator-lock"
	var defaultDesiredReplicas int32 = 1
	var clusterRoles = []string{
		"managed-velero-operator",
	}
	var clusterRoleBindings = []string{
		"managed-velero-operator",
		"velero",
	}
	h := helper.New()
	checkConfigMapLockfile(h, operatorNamespace, operatorLockFile)
	checkDeployment(h, operatorNamespace, operatorName, defaultDesiredReplicas)
	checkDeployment(h, operatorNamespace, "velero", defaultDesiredReplicas)
	checkClusterRoles(h, clusterRoles)
	checkClusterRoleBindings(h, clusterRoleBindings)
	checkRoleBindings(h,
		operatorNamespace,
		[]string{"managed-velero-operator"})
	checkVeleroBackups(h)

	ginkgo.It("Access should be forbidden to edit Backups", func() {
		h.SetServiceAccount("system:serviceaccount:%s:dedicated-admin-project")

		backup := velerov1.Backup{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dedicated-admin-backup-test",
			},
		}
		_, err := h.Velero().VeleroV1().Backups(h.CurrentProject()).Create(context.TODO(), &backup, metav1.CreateOptions{})
		Expect(apierrors.IsForbidden(err)).To(BeTrue())
	})

	ginkgo.It("Access should be forbidden to edit DeleteBackupRequests", func() {
		h.SetServiceAccount("system:serviceaccount:%s:dedicated-admin-project")

		delete_backup_request := velerov1.DeleteBackupRequest{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dedicated-admin-delete-backup-request-test",
			},
		}
		_, err := h.Velero().VeleroV1().DeleteBackupRequests(h.CurrentProject()).Create(context.TODO(), &delete_backup_request, metav1.CreateOptions{})
		Expect(apierrors.IsForbidden(err)).To(BeTrue())
	})

	ginkgo.It("Access should be forbidden to edit BackupStorageLocations", func() {
		h.SetServiceAccount("system:serviceaccount:%s:dedicated-admin-project")

		backup_storage_location := velerov1.BackupStorageLocation{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dedicated-admin-backup-storage-location-test",
			},
		}
		_, err := h.Velero().VeleroV1().BackupStorageLocations(h.CurrentProject()).Create(context.TODO(), &backup_storage_location, metav1.CreateOptions{})
		Expect(apierrors.IsForbidden(err)).To(BeTrue())
	})

	ginkgo.It("Access should be forbidden to edit DownloadRequests", func() {
		h.SetServiceAccount("system:serviceaccount:%s:dedicated-admin-project")

		download_request := velerov1.DownloadRequest{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dedicated-admin-download-request-test",
			},
		}
		_, err := h.Velero().VeleroV1().DownloadRequests(h.CurrentProject()).Create(context.TODO(), &download_request, metav1.CreateOptions{})
		Expect(apierrors.IsForbidden(err)).To(BeTrue())
	})

	ginkgo.It("Access should be forbidden to edit PodVolumeBackups", func() {
		h.SetServiceAccount("system:serviceaccount:%s:dedicated-admin-project")

		pod_volume_backup := velerov1.PodVolumeBackup{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dedicated-admin-pod-volume-backup-test",
			},
		}
		_, err := h.Velero().VeleroV1().PodVolumeBackups(h.CurrentProject()).Create(context.TODO(), &pod_volume_backup, metav1.CreateOptions{})
		Expect(apierrors.IsForbidden(err)).To(BeTrue())
	})

	ginkgo.It("Access should be forbidden to edit PodVolumeRestores", func() {
		h.SetServiceAccount("system:serviceaccount:%s:dedicated-admin-project")

		pod_volume_restore := velerov1.PodVolumeRestore{
			ObjectMeta: metav1.ObjectMeta{
				Name: "dedicated-admin-pod-volume-restore-test",
			},
		}
		_, err := h.Velero().VeleroV1().PodVolumeRestores(h.CurrentProject()).Create(context.TODO(), &pod_volume_restore, metav1.CreateOptions{})
		Expect(apierrors.IsForbidden(err)).To(BeTrue())
	})

})

func checkVeleroBackups(h *helper.H) {
	ginkgo.Context("velero", func() {
		ginkgo.It("backups should be complete", func() {
			err := wait.PollImmediate(5*time.Second, 5*time.Minute, func() (bool, error) {
				us, err := h.Dynamic().Resource(schema.GroupVersionResource{
					Group:    "velero.io",
					Version:  "v1",
					Resource: "backups"}).Namespace(operatorNamespace).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					return false, fmt.Errorf("Error getting backups: %s", err.Error())
				}
				for _, item := range us.Items {
					var backup velerov1.Backup
					err = runtime.DefaultUnstructuredConverter.
						FromUnstructured(item.UnstructuredContent(), &backup)
					if err != nil {
						return false, fmt.Errorf("Error casting object: %s", err.Error())
					}
					if backup.Status.Phase == velerov1.BackupPhaseFailed {
						return false, fmt.Errorf("Backup failed for %s", backup.Name)
					}
					if backup.Status.Phase != velerov1.BackupPhaseCompleted {
						return false, nil
					}
				}
				return true, nil
			})
			Expect(err).ToNot(HaveOccurred(), "Backups failed to complete.")
		})
	})
}
