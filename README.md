# Library-management


1. С первым делом нужно поднимать **library-management**

```docker compose up --build app```

2. После успешного запуска **one_edu** запускаем второй микроссервис по [ссылке](https://github.com/jumagaliev1/transactions-go)

```docker compose up --build app```


При возникновение ошибок с база данных попробуем 

```docker compose up```

---
Main endpoints:
аренда книги:
- клиент может взять книгу в аренду за опр сумму 
![image](https://user-images.githubusercontent.com/71185943/233068625-50315d43-86ff-44e9-80cb-6bcf02c5ddf4.png)
- список книг которые сейчас у клиентов и суммарный доход по каждой 
![image](https://user-images.githubusercontent.com/71185943/233068728-44693686-33b6-4700-9dab-589afbd9f237.png)



