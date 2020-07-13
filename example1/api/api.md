<style>
span {padding: 8px;}
.method {
	border-radius: 6px 0px 0px 6px;
	color: #FFF;
	min-width: 80px;
}
.url {
	background-color: #D5D5D5;
	color: #000;
	border-radius: 0px 6px 6px 0px;
}
.get {background-color: #61affe}
.post {background-color: #49cc90}
.delete {background-color: #f93e3e}
.put{background-color: #fca130}
.patch{background-color: #00b3b3}
a {text-decoration: none;}
</style>

# Service Group API
v 0.0.1


<span class="get method"><b>GET</b></span><span class="url"><b>/</b> - получение списка людей</span>

>**type:**application/json

>**Responses:**
>***Code:*** **200** - всё хорошо
>**Body**:

>		{
>			"int":
>				{
>					"id": int,
>					"name": string,
>					"surname": string
>				},
>			...
>		}
>***Code:*** **400** - плохой неверный запрос, ошибка в json

---

<span class="get method"><b>GET</b></span><span class="url"><b>/person?id=</b> - найти человека по id</span>

>**type:**application/json
> **Params:**
 
>		id : int

>**Responses:**
>***Code:*** **200** - удачный вход в учетную запись
>**Body**:

>		{
>			"id": int,
>			"name": string,
>			"surname": string
>		}

>***Code:*** **400** - плохой неверный запрос, ошибка в json

>***Code:*** **404** - человека с таким id не существует

-----

<span class="post method"><b>POST</b></span><span class="url"><b>/person</b> - создание нового человка</span>


>**type:**application/json
>**Reqest:**

>		{
>			"name": string,
>			"surname": string
>		}
>	
>**Responses:**
>***Code:*** **202** - запись создана
>**Body**:

>		{
>			"int":
>				{
>					"id": int,
>					"name": string,
>					"surname": string
>				},
>			...
>		}


>***Code:*** **400** - плохой неверный запрос, ошибка в json

---

<span class="put method"><b>PUT</b></span><span class="url"><b>/person</b> -обновление записи </span>


>**type:**application/json
>**Reqest:**

>		{
>			"id" : int
>			"name": string,
>			"surname": string
>		}
>	
>**Responses:**
>***Code:*** **200** - запись обновленна
>**Body**:

>		{
>			"int":
>				{
>					"id": int,
>					"name": string,
>					"surname": string
>				},
>			...
>		}


>***Code:*** **400** - плохой неверный запрос, ошибка в json

----
	
<span class="delete method"><b>DELETE</b></span><span class="url"><b>/person</b> - удаление записи</span>

>**type:**application/json
>**Reqest:**

>		{
>			"id": int,
>		}
>	
>**Responses:**
>***Code:*** **200** - запись удалена
>**Body**:

>		{
>			"int":
>				{
>					"id": int,
>					"name": string,
>					"surname": string
>				},
>			...
>		}


>***Code:*** **400** - плохой неверный запрос, ошибка в json

---

