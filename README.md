# Project 3 - Golang Hacktiv8 X Kampus Merdeka

Scalable Web Services with Golang

Kelompok : 
- Andri Nur Hidayatulloh
- Andri Kuwito

### List library : 
- [Gin Gionic](https://github.com/gin-gonic/gin) - Web Framework
- [Gorm](https://gorm.io) - GORM
- [Go Validator](https://github.com/asaskevich/govalidator) - Go Validator


### Cara Penggunaan : 

* ##### URL : https://morning-taiga-16831.herokuapp.com/


* #### Endpoint List : 
    * ##### User : 
        * #### Register
            
            [POST]```https://murmuring-savannah-72759.herokuapp.com/users/register```
            
            body :

            ```json
            {
                "full_name": "Andri Nur H",
                "password": "Bismillah",
                "email": "andriandri@gmail.com"
            }
            ```

            response :
            ```json
            {
                "status": "success",
                "data": {
                    "id": 7,
                    "full_name": "Andri Nur H",
                    "email": "andriandri@gmail.com",
                    "created_at": "2022-01-22T08:14:54.646Z"
                }
            }
            ```
            
        * #### Login
            
            [POST]```https://murmuring-savannah-72759.herokuapp.com/users/login```
            
            body :

            ```json
            {
                "email": "andri@gmail.com",
                "password": "Bismillah"
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjozfQ.ZtB0-hBckTjyBjDyQNko06sIuPWAvp8Ja0rNmVVOs9M"
                }
            }
            ```
            
        * #### Update Account
            
            [DELETE]```https://murmuring-savannah-72759.herokuapp.com/users/update-account```
            
            headers :

            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :

            ```json
            {
                "full_name": "Andri Hidayatulloh",
                "email": "andri@gmail.com"
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 3,
                    "full_name": "Andri Hidayatulloh",
                    "email": "andri@gmail.com",
                    "updated_at": "2022-01-22T08:21:27Z"
                }
            }
            ```
            
        * #### Delete Account
            
            [PUT]```https://murmuring-savannah-72759.herokuapp.com/users/delete-account```
            
            headers :

            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": "Your account has been successfully deleted!"
            }
            ```
            
    * ##### Category : 
        * #### Add Category
            
            [POST]```https://murmuring-savannah-72759.herokuapp.com/categories```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "type": "Pekerjaan Berat"
            }
            ```

            response :
            ```json
            {
                "status": "success",
                "data": {
                    "id": 7,
                    "types": "Pekerjaan Berat",
                    "created_at": "2022-01-22T17:40:14.503+07:00"
                }
            }
            ```
            
        * #### Get Category
            
            [GET]```https://murmuring-savannah-72759.herokuapp.com/categories```
            
            headers :

            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```

            response :
            ```json
            {
            "status": "success",
            "data": [
                {
                    "id": 3,
                    "type": "Pekerjaan Rumah",
                    "created_at": "2022-01-18T21:34:57+07:00",
                    "updated_at": "2022-01-18T21:43:15+07:00"
                },
                {
                    "id": 4,
                    "type": "Pekerjaan Kantor",
                    "created_at": "2022-01-18T21:36:28+07:00",
                    "updated_at": "2022-01-18T21:36:28+07:00"
                },  
            }
            ```
            
        * #### Update Category
            
            [PATCH]```https://murmuring-savannah-72759.herokuapp.com/categories/:id```
            
            headers :

            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "type": "Pekerjaan Backend"
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 3,
                    "types": "Pekerjaan Backend",
                    "updated_at": "2022-01-22T17:43:32+07:00"
                }
            }
            ```
            
        * #### Delete Category
            
            [DELETE]```https://murmuring-savannah-72759.herokuapp.com/categories/:id```
            
            headers :

            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": "Category deleted"
            }
            ```
            
    * ##### Task : 
        * #### Add Task
            
            [POST]```https://murmuring-savannah-72759.herokuapp.com/tasks```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "title": "Nyuci Mobil",
                "description": "Nyuci mobik tiap sore",
                "category_id": 2
            }
            ```

            response :
            ```json
            {
                "status": "success",
                "data": {
                    "id": 6,
                    "title": "Nyuci Mobil",
                    "description": "Nyuci mobik tiap sore",
                    "status": false,
                    "user_id": 3,
                    "category_id": 2,
                    "created_at": "2022-01-22T17:39:39.923+07:00"
                }
            }
            ```
            
        * #### GET Task
            
            [GET]```https://murmuring-savannah-72759.herokuapp.com/tasks```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```

            response :
            ```json
            {
                "status": "success",
                "data": [
                    {
                        "id": 2,
                        "title": "Nyuci Baju",
                        "description": "Nyuci baju tiap minggu",
                        "status": false,
                        "user_id": 4,
                        "category_id": 3,
                        "created_at": "2022-01-16T17:48:04+07:00",
                        "updated_at": "2022-01-21T22:47:31+07:00",
                        "deleted_at": "0001-01-01T00:00:00Z"
                    }
                ]
            }
            ```
        
        * #### Update Task
            
            [PUT]```https://murmuring-savannah-72759.herokuapp.com/tasks/:id```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "title": "Nyuci Truck",
                "description": "deksirpsi test"
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 4,
                    "title": "Nyuci Truck",
                    "description": "deksirpsi test",
                    "status": false,
                    "user_id": 3,
                    "category_id": 2,
                    "updated_at": "2022-01-22T17:51:48+07:00"
                }
            }
            ```
        
        * #### Update Status Task
            
            [PATCH]```https://murmuring-savannah-72759.herokuapp.com/tasks/update-status/:id```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 4,
                    "title": "Nyuci Truck",
                    "description": "deksirpsi test",
                    "status": false,
                    "user_id": 3,
                    "category_id": 2,
                    "updated_at": "2022-01-22T17:54:15+07:00"
                }
            }
            ```
            
        * #### Update Category Task
            
            [PATCH]```https://murmuring-savannah-72759.herokuapp.com/tasks/update-category/:id```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            body :
            ```json
            {
                "category_id": 3
            }
            ```

            response :
            ```json
            {
                "status": "ok",
                "data": {
                    "id": 2,
                    "title": "Nyuci Baju",
                    "description": "Nyuci baju tiap minggu",
                    "status": false,
                    "user_id": 4,
                    "category_id": 3,
                    "updated_at": "2022-01-21T22:47:31+07:00"
                }
            }
            ```
        * #### Delete Task
            
            [DELETE]```https://murmuring-savannah-72759.herokuapp.com/tasks/:id```
            
            headers :
            ```
            {
                "Authorization": "Bearer {{token}}"
            }
            ```
            
            response :
            ```json
            {
                "status": "ok",
                "data": "Task has been successfully deleted"
            }
            ```
           
        