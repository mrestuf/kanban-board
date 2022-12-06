# kanban-board

## directory structure 
```
   .
    ├── common                   # Response JSON data
    ├── cmd                      # Include main.go
    ├── config                   # Configure Database Connection 
    ├── httpserver               # Include controllers, repositories, services and router.go
    ├── .env.example             # copying secret variable to connect db from .env
    ├── .gitignore               # hiding .env
    ├── go.mod                 
    ├── go.sum                   
    ├── LICENSE
    └── 
```
### run kanban-board
```
first must open query on postgres and run query CREATE TYPE role AS ENUM ('member', 'admin');
then
go run cmd/main.go
```

## list endpoint

### Register : Registration user account
```
POST https://kanban-board-production-a581.up.railway.app/v1/users/register
```

### Login : Login user account
```
POST https://kanban-board-production-a581.up.railway.app/v1/users/login
```
### Update-Account : Edit and Update User account
```
PUT https://kanban-board-production-a581.up.railway.app/v1/users/update-account
```

### Delete-Account : Delete User account from list
```
DELETE https://kanban-board-production-a581.up.railway.app/v1/users/delete-account
```
### CreateCategory : Create Category (only for admin) 
```
POST https://kanban-board-production-a581.up.railway.app/v1/categories
```

### GetCategories : Display all categories with tasks that included for each category
```
GET https://kanban-board-production-a581.up.railway.app/v1/categories
```

### UpdateCategory : Edit and Update Category Type (only for admin)
```
PATCH https://kanban-board-production-a581.up.railway.app/v1/categories/:categoryId
```

### DeleteCategory : Delete Category from list (only for admin)
```
DELETE https://kanban-board-production-a581.up.railway.app/v1/categories/:categoryId
```
### GetTasks : Display all tasks list  
```
GET https://kanban-board-production-a581.up.railway.app/v1/tasks
```

### CreateTask : Create task from specific user member
```
POST https://kanban-board-production-a581.up.railway.app/v1/tasks
```

### UpdateTask : Edit and Update Task title and description
```
PUT https://kanban-board-production-a581.up.railway.app/v1/tasks/:tasksId
```
### UpdateTask-Status : Edit and Update Task Status
```
PATCH https://kanban-board-production-a581.up.railway.app/v1/tasks/update-status/:tasksId
```
### UpdateTask-Category : Edit and Update Task Category
```
PATCH https://kanban-board-production-a581.up.railway.app/v1/tasks/update-category/:tasksId
```

### DeleteTask : Delete Task from list
```
DELETE https://kanban-board-production-a581.up.railway.app/v1/tasks/:tasksId
```

### Jobdesk member

- MAULA IZZA AZIZI (GLNG-KS04-020) : 
   -  User EndPoint     : - Login
                          - Update
                          - Delete
   -  Category EndPoint : - GetCategories


- HEZKYA NATANAEL RAMLI (GLNG-KS04-008) : 
   -  All Tasks Endpoint
   -  Category EndPoint : - GetCategories
   -  Mocking
   -  Unit Test
   -  Deployment

- MUHAMAD RESTU FADILLAH (GLNG-KS04-002) : 
   -  User EndPoint     : - Register
   -  Category EndPoint : - Create Category
                          - Update Category
                          - Delete Category 
   -  Postman Collection & README.md
