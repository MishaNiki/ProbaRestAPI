Example1 API

**[GET]: /**  - получение списка людей

**type:**application/json

**Responses:**
***Code:*** **200** - всё хорошо
**Body**:

```
{
	"int":
	{
		"id": int,
		"name": string,
		"surname": string
	},
	...
}
```
***Code:*** **400** - плохой неверный запрос, ошибка в json

---

**[GET]: /person?id=** - найти человека по id

**type:**application/json
 **Params:**

	id : int

**Responses:**

***Code:*** **200** - удачный вход в учетную запись
**Body**:

```
{
    "id": int,
    "name": string,
    "surname": string
}
```
***Code:*** **400** - плохой неверный запрос, ошибка в json

***Code:*** **404** - человека с таким id не существует

-----

**[POST]: /person** - создание нового человка

**type:**application/json
**Reqest:**

```
{
    "name": string,
    "surname": string
}
```

**Responses:**
***Code:*** **202** - запись создана
**Body**:

```
{
    "int":
    {
        "id": int,
        "name": string,
        "surname": string
    },
    ...
}
```

***Code:*** **400** - плохой неверный запрос, ошибка в json

---

**[PUT]: /person** -обновление записи


**type:**application/json
**Reqest:**

```
{
    "id" : int
    "name": string,
    "surname": string
}
```

**Responses:**
***Code:*** **200** - запись обновленна
**Body**:

```
{
    "int":
    {
        "id": int,
        "name": string,
        "surname": string
    },
    ...
}
```

***Code:*** **400** - плохой неверный запрос, ошибка в json

----

**[DELETE]: /person** - удаление записи

**type:**application/json
**Reqest:**

```
{
	"id": int,
}
```

**Responses:**
***Code:*** **200** - запись удалена
**Body**:

```
{
    "int":
    {
        "id": int,
        "name": string,
        "surname": string
    },
    ...
}
```

***Code:*** **400** - плохой неверный запрос, ошибка в json

---

