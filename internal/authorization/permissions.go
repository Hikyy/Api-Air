package authorization

type Permission struct {
	ID          int
	Name        string
	Description string
}

func GetAllPermissions() ([]Permission, error) {
	// Implémentation pour récupérer tous les rôles depuis la base de données
}

func GetPermission() (Permission, error) {
	// Implémentation pour récupérer tous les rôles depuis la base de données
}

func GetPermissionByID(id int) (Permission, error) {
	// Implémentation pour récupérer un rôle par son ID depuis la base de données
}
