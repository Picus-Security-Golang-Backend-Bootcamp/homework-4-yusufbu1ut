# Homework-4 / yusufbu1ut
## About Work 

## Books

Getting all books <br/>
`GET localhost:8090/books` <br/>
`GET localhost:8090/books/`<br/>
Returns all books with its authors

Getting books with id <br/>
`GET localhost:8090/books/:id` <br/>
`GET localhost:8090/books/id=` <br/>
Returns book with its authors

Getting books with search <br/>
`GET localhost:8090/books/search=` <br/>
Returns books with authors which contains search parameter in book name

Deleting books with id <br/>
`DELETE localhost:8090/books/:id` <br/>
`DELETE localhost:8090/books/id=` <br/>
Deletes connected elements on db

Creating books <br/>
`POST localhost:8090/books/create` <br/>
Creates book on db with request body

 {<br/>
        "name": "Fabl",<br/>
        "pages": "45",<br/>
        "publisher": "Scholastic",<br/>
        "code": 671760,<br/>
        "amount": 3,<br/>
        "isbn": 8880439554893,<br/>
        "price": 10.79,<br/>
        "authors": [ <br/>
            {<br/>
                "author":"Lafontane"<br/>
            }
        ]
}<br/>
ps: ID,name,amount,isbn vals required. Can be added without author, if there is author author name is required.

Updating books <br/>
`PATCH localhost:8090/books/update` <br/>
Updates book on db with request body

{<br/>
        "ID": 1,<br/>
        "CreatedAt": "2022-03-26T02:36:58.891559+03:00",<br/>
        "UpdatedAt": "2022-03-26T03:18:37.50501+03:00",<br/>
        "DeletedAt": null,<br/>
        "name": "Harry Potter and the Chamber of Secrets (Harry Potter  #2)",<br/>
        "pages": "352",<br/>
        "publisher": "Scholastic",<br/>
        "code": 671760,<br/>
        "amount": 3,<br/>
        "isbn": 9780439554893,<br/>
        "price": 122.79,<br/>
        "authors": [ ]<br/>
}<br/>
ps: Changeable variables are name,pages,publisher,amount,price. If add any author with requested field author program arranges book and authors connections.<br/>
ps: ID,name,amount,isbn vals required. if there is author author name is required. In here author phase to connect books with authors.

Buyying books <br/>
`POST localhost:8090/books/:id/count=` <br/>
Buys book and updates amount on db 

## Authors

Getting all authors <br/>
`GET localhost:8090/authors` <br/>
`GET localhost:8090/authors/`<br/>
Returns all authors with its books

Getting authors with id <br/>
`GET localhost:8090/authors/:id` <br/>
`GET localhost:8090/authors/id=` <br/>
Returns author with its books

Getting authors with search <br/>
`GET localhost:8090/authors/search=` <br/>
Returns authors with books which contains search parameter in authors name

Deleting authors with id <br/>
`DELETE localhost:8090/authors/:id` <br/>
`DELETE localhost:8090/authors/id=` <br/>
Deletes connected elements on db

Creating authors <br/>
`POST localhost:8090/authors/create` <br/>
Creates author on db with request body

Updating authors <br/>
`PATCH localhost:8090/authors/update` <br/>
Updates author on db with request body

ps: In Author part processes are same, just values different.

## Database

In this work, I used PostgreSQL.

### books

ID         
CreatedAt   
UpdatedAt   
DeletedAt   
Name
Pages       
Publisher   
StockCode     
StockAmount       
ISBN    (unique count and foreing key with referance book_authors BookID)    
Price       

### book_authors

ID         
CreatedAt   
UpdatedAt   
DeletedAt   
BookID   (foreing key with referance books ISBN)  
AuthorID    (foreing key with referance authors ID)  

### authors

ID         
CreatedAt   
UpdatedAt   
DeletedAt   
NameSurname     
Age           
