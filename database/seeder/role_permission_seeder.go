package seeder

import (
	"TixTrain/app/model"
	"TixTrain/pkg"
	"log"
)

func SeedRolePermissions() error {
	log.Println("Seeding role permissions...")

	// Fetch all roles and permissions
	var roles []model.Role
	var permissions []model.Permission

	if err := pkg.DB.Find(&roles).Error; err != nil {
		return err
	}

	if err := pkg.DB.Find(&permissions).Error; err != nil {
		return err
	}

	// Create a map for easier lookup
	permissionMap := make(map[string]*model.Permission)
	roleMap := make(map[string]*model.Role)

	for i := range permissions {
		permissionMap[permissions[i].Name] = &permissions[i]
	}

	for i := range roles {
		roleMap[roles[i].Name] = &roles[i]
	}

	// Assign permissions to roles based on api.go routes
	if adminRole, ok := roleMap["admin"]; ok {
		// Admin routes: /admin/stats, /admin/tickets, /admin/tickets/all, /admin/stations
		permissionNames := []string{
			"view_admin_stats",
			"view_admin_tickets",
			"view_admin_tickets_all",
			"manage_admin_stations",
		}
		for _, permName := range permissionNames {
			if perm, ok := permissionMap[permName]; ok {
				rp := model.RolePermission{
					RoleID:       adminRole.ID,
					PermissionID: perm.ID,
				}
				pkg.DB.Create(&rp)
			}
		}
	}

	if staffRole, ok := roleMap["staff"]; ok {
		// Staff routes: /staff/tickets
		permissionNames := []string{
			"view_staff_tickets",
		}
		for _, permName := range permissionNames {
			if perm, ok := permissionMap[permName]; ok {
				rp := model.RolePermission{
					RoleID:       staffRole.ID,
					PermissionID: perm.ID,
				}
				pkg.DB.Create(&rp)
			}
		}
	}

	if passengerRole, ok := roleMap["passenger"]; ok {
		// Passenger routes: /passenger/tickets
		permissionNames := []string{
			"view_passenger_tickets",
		}
		for _, permName := range permissionNames {
			if perm, ok := permissionMap[permName]; ok {
				rp := model.RolePermission{
					RoleID:       passengerRole.ID,
					PermissionID: perm.ID,
				}
				pkg.DB.Create(&rp)
			}
		}
	}

	log.Println("Role permissions seeding completed")
	return nil
}
