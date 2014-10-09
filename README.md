#Rest API List(alpha)


##用户
1. 登陆 /users/login  [get] phone=;password=
2. 注册 /users [post] phone=;password=
3. 修改个人信息 /users/ID [put] phone;password;email;gender;birthday;name;nick;head;
4. 关注 /ul [post] following;
5. 取消关注 /ul/ID[delete]
5. 关注列表 /ul/following [get]
6. 粉丝列表 /ul/follower [get]
7. 注销 /users/logout [get]

##图片
1. 发布照片 /photos [post]photo;title
2. 我的照片列表 /photos [get]offset;
2. 喜欢 /likes/ID [post] photo_id;
3. 取消喜欢 /likes/ID [delete]
3. 评论 /comment[post]photo_id;content;
4. 删除评论 /comment/ID [delete]


##专题文章	

1. 创建专题 /articles [post] title;content
2. 专题列表 /articles [get]
3. 更新专题 /articles/ID [put]title;content;
4. 专题详情 /articles/ID [get]

##统一返回格式Json
* Err:int (0为正常，1为错误)
* Data:（为数组或者对象数据）
* Msg:（string,提示信息）

```
{
    "Err": 0,
    "Data": {
        "Id": 16,
        "CreatedAt": "2014-10-09T17:00:42.761450016+08:00",
        "UpdatedAt": "2014-10-09T17:00:42.76145004+08:00",
        "TargetId": {
            "Id": 43,
            "Title": "ttttt",
            "Path": "static/uploads/photos/2014-10-08/be/bea280464513b4e4ced845d0ea89cd0d.jpg",
            "CreatedAt": "2014-10-08T12:46:32+08:00",
            "UpdatedAt": "2014-10-08T12:46:32+08:00",
            "UserId": {
                "Id": 29,
                "Email": "forevervmax@gmail.com",
                "Password": "123",
                "Gender": 0,
                "Phone": "18611358272",
                "Birthday": "",
                "CreatedAt": "2014-10-05T15:23:27+08:00",
                "UpdatedAt": "2014-10-05T15:23:27+08:00",
                "Name": "tttt",
                "Following": 1,
                "Follower": 0,
                "Head": ""
            },
            "Likes": 0
        },
        "UserId": {
            "Id": 29,
            "Email": "forevervmax@gmail.com",
            "Password": "123",
            "Gender": 0,
            "Phone": "18611358272",
            "Birthday": "",
            "CreatedAt": "2014-10-05T15:23:27+08:00",
            "UpdatedAt": "2014-10-05T15:23:27+08:00",
            "Name": "tttt",
            "Following": 1,
            "Follower": 0,
            "Head": ""
        }
    },
    "Msg": ""
}
```
```
{
    "Err": 0,
    "Data": [
        {
            "CreatedAt": "2014-10-09T17:05:18+08:00",
            "Following__Id": 36,
            "Following__Name": "fffff"
        },
        {
            "CreatedAt": "2014-10-09T17:05:16+08:00",
            "Following__Id": 35,
            "Following__Name": "1231\n"
        },
        {
            "CreatedAt": "2014-10-09T16:53:21+08:00",
            "Following__Id": 32,
            "Following__Name": "afaf"
        }
    ],
    "Msg": ""
}
```


```
{
    "Err": 1,
    "Data": "",
    "Msg": "Error 1062: Duplicate entry 'forevervmax@gmail.com' for key 'email'"
}
```