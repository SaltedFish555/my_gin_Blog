// 登录函数
async function login() {
    try {
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        const requestBody = {
            username: username,
            password: password,
        };

        const response = await fetch('http://8.130.162.251:3389/api/v1/login', {
            method: 'POST',
            body: JSON.stringify(requestBody)
        });

        const contentType = response.headers.get('content-type');
        if (!contentType || !contentType.includes('application/json')) {
            throw new Error('服务器返回的不是有效的 JSON 数据');
        }

        const data = await response.json();

        console.log('登录结果:', data);

        if (data.status === 200) {
            console.log('登录成功，Token:', data.token);
            window.location.href = 'index.html';
            sessionStorage.setItem('token', data.token);
        } else {
            console.error('登录失败:', data.message);
            // 处理登录失败的逻辑，例如显示错误消息
        }
    } catch (error) {
        console.error('登录失败:', error.message);
        // 处理错误，例如显示错误消息
    }
}


// 退出登录函数
function logout() {
    sessionStorage.removeItem('token');
    changeLink();
    showCustomToast('退出成功!')
}


// 获取用户列表
function getUsers() {
    // 构造 GET 请求的地址
    const apiUrl = 'http://8.130.162.251:3389/api/v1/users?pagesize=2&pagenum=1';  // 替换成你的实际接口地址

    // 发送 GET 请求
    // fetch(apiUrl)
    fetch(apiUrl, {
        method: 'GET',
        headers: {
            // 添加 CORS 相关 headers
            // 暂时好像都不需要
            // 'Origin': 'http://localhost:3389'
            // 其他可能需要的 headers...
        },
    })

        .then(response => {
            // 检查响应是否为 JSON 格式
            const contentType = response.headers.get('content-type');
            if (!contentType || !contentType.includes('application/json')) {
                throw new Error('服务器返回的不是有效的 JSON 数据');
            }

            return response.json();

        })
        .then(data => {

            // 处理接口返回的数据
            console.log('接口返回的数据:', data);

        })
        .catch(error => {
            // 处理错误
            console.error('请求失败:', error.message);

            // 根据实际情况进行错误处理，例如显示错误消息
        });
}

// 存储分类字典，键为分类名，值为分类的id(即主键)
// var categories = { 'Golang': 1 };
// 获取分类列表
function fetchCategories() {
    // 构造 GET 请求的地址
    const apiUrl = 'http://8.130.162.251:3389/api/v1/categories?pagesize=10&pagenum=1';  // 替换成你的实际接口地址

    // 发送 GET 请求
    return fetch(apiUrl, {
        method: 'GET',
        headers: {
            // 添加 CORS 相关 headers
            // 暂时好像都不需要
            // 'Origin': 'http://localhost:3389'
            // 其他可能需要的 headers...
        },
    })
        .then(response => {
            // 检查响应是否为 JSON 格式
            const contentType = response.headers.get('content-type');
            if (!contentType || !contentType.includes('application/json')) {
                throw new Error('服务器返回的不是有效的 JSON 数据');
            }

            return response.json();
        })
        .then(data => {
            // 返回处理后的数据
            return data.data.reduce((acc, category) => {
                acc[category.name] = category.id;
                return acc;
            }, {});
        })
        .catch(error => {
            // 处理错误
            console.error('请求失败:', error.message);

            // 根据实际情况进行错误处理，例如显示错误消息
            throw error;  // 抛出错误以便上层处理
        });
}

// 显示分类列表在页面上
function displayCategories(categories) {
    console.log('接口返回的数据:', categories);

    // 在这里添加处理分类的逻辑，例如展示在页面上
    for (const name in categories) {
        const listItem = document.createElement('li');
        const link = document.createElement('a');
        link.href = `cate_articles.html?cname=${encodeURIComponent(name)}&cid=${encodeURIComponent(categories[name])}`;
        link.textContent = name;
        listItem.appendChild(link);
        categoryList.appendChild(listItem);
    }

}





function getCategoryFromUrl() {


    const searchParams = new URLSearchParams(window.location.search);
    const cname = searchParams.get('cname');
    const cid = searchParams.get('cid');
    console.log(cname, cid)
    document.getElementById('pageTitle').innerText = cname;

    return { cname, cid };
}





// Assuming you have a function to fetch articles for a specific category
// Replace fetchArticles() with your actual function
function fetchCateArticles({ cname, cid }) {
    // 构造 GET 请求的地址
    // 获取某个分类下的文章列表
    const apiUrl = `http://8.130.162.251:3389/api/v1/article/clist/${cid}?pagesize=10&pagenum=1`; // 替换成你的实际接口地址
    console.log(apiUrl)
    // 发送 GET 请求
    // fetch(apiUrl)
    fetch(apiUrl, {
        method: 'GET',
        headers: {
            // 添加 CORS 相关 headers
            // 暂时好像都不需要
            // 'Origin': 'http://localhostsetUserArtLink:3389'
            // 其他可能需要的 headers...
        },
    })

        .then(response => {
            // 检查响应是否为 JSON 格式
            const contentType = response.headers.get('content-type');
            if (!contentType || !contentType.includes('application/json')) {
                throw new Error('服务器返回的不是有效的 JSON 数据');
            }

            return response.json();

        })
        .then(data => {
            // 处理接口返回的数据
            console.log('接口返回的数据:', data);

            // 提取每篇文章的标题和id
            const articles = data.data.map(article => ({
                title: article.title,
                id: article.ID,
            }));

            // 在这里添加处理文章标题和id的逻辑，例如展示在页面上
            console.log('文章列表:', articles);

            const articleList = document.getElementById('articleList');
            articles.forEach(article => {
                const listItem = document.createElement('li');
                const link = document.createElement('a');
                link.href = `article_content.html?cname=${encodeURIComponent(cname)}&cid=${encodeURIComponent(cid)}&title=${encodeURIComponent(article.title)}&id=${encodeURIComponent(article.id)}`;
                link.textContent = article.title;
                listItem.appendChild(link);
                articleList.appendChild(listItem);
            });

        })
        .catch(error => {
            // 处理错误
            console.error('请求失败:', error.message);

            // 根据实际情况进行错误处理，例如显示错误消息
        });

}

function getTitleFromUrl() {
    // Get the current URL
    const url = window.location.href;

    // Output the full URL to check if it's correct
    // Extract the title parameter from the URL
    const urlParams = new URLSearchParams(window.location.search);

    // const searchParams = new URLSearchParams(window.location.search);
    // const category = searchParams.get('category');

    // Output the URL parameters to check if 'title' is present
    console.log('URL Parameters:', urlParams);

    const title = urlParams.get('title');
    const id = urlParams.get('id');


    // Output the extracted title to check its value
    console.log('Title:', title);
    console.log('id:', id);


    return { title, id };
}


// Assuming you have a function to fetch article content by title and id
// Replace fetchArticleContent() with your actual function
function fetchArticleContent({ title, id }) {
    // 构造 GET 请求的地址
    const apiUrl = `http://8.130.162.251:3389/api/v1/article/info/${id}`; // 替换成你的实际接口地址
    console.log('apiurl:', apiUrl)
    console.log('Title:', title);
    console.log('id:', id);

    // 发送 GET 请求
    fetch(apiUrl, {
        method: 'GET',
        headers: {
            // 添加 CORS 相关 headers
            // 暂时好像都不需要
            // 'Origin': 'http://localhost:3389'
            // 其他可能需要的 headers...
        },
    })
        .then(response => {
            // 检查响应是否为 JSON 格式
            const contentType = response.headers.get('content-type');
            if (!contentType || !contentType.includes('application/json')) {
                throw new Error('服务器返回的不是有效的 JSON 数据');
            }

            return response.json();

        })
        .then(data => {
            // 处理接口返回的数据
            console.log('接口返回的数据:', data);

            // 提取文章内容
            const articleTitle = document.getElementById('articleTitle');
            const articleContentContainer = document.getElementById('articleContent');

            // Display article title
            articleTitle.textContent = data.data.title;

            // Display article desc
            const descParagraph = document.createElement('p');
            descParagraph.textContent = `描述:\n ${data.data.desc}`;
            articleContentContainer.appendChild(descParagraph);

            // Display article content
            const contentParagraph = document.createElement('p');
            contentParagraph.textContent = `内容:\n ${data.data.content}`;
            articleContentContainer.appendChild(contentParagraph);

        })
        .catch(error => {
            // 处理错误
            console.error('请求失败:', error.message);

            // 根据实际情况进行错误处理，例如显示错误消息
        });
}


// 注册函数
async function register() {
    // 获取表单中的值
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;

    // 简单的输入验证
    if (!username || !password || !confirmPassword) {
        showCustomToast('请填写所有字段');
        return;
    }

    // 长度验证
    if (username.length < 4 || username.length > 12) {
        showCustomToast('用户名长度应在4到12个字符之间');
        return;
    }

    if (password.length < 6 || password.length > 20) {
        showCustomToast('密码长度应在6到20个字符之间');
        return;
    }

    if (password !== confirmPassword) {
        showCustomToast('密码和确认密码不一致');
        return;
    }

    // 构造请求体
    const requestBody = {
        username: username,
        password: password,
        role: 2,
    };
    try {
        const response = await fetch('http://8.130.162.251:3389/api/v1/user/add', {
            method: 'POST',
            headers: {
            },
            body: JSON.stringify(requestBody),
        });

        // 检查响应是否为 JSON 格式
        const contentType = response.headers.get('content-type');
        if (!contentType || !contentType.includes('application/json')) {
            throw new Error('服务器返回的不是有效的 JSON 数据');
        }

        const data = await response.json();

        console.log('注册结果:', data);

        if (data.status === 200) {
            console.log('注册成功，Token:', data.token);
            showCustomToast('注册成功，等待页面跳转中 ————')
            // 注册成功后自动登录
            await login();
            window.location.href = 'index.html';
            // 可以添加页面跳转等其他逻辑
        } else {
            console.error('注册失败:', data.message);
            showCustomToast(`注册失败 ${data.message}`);
            // 根据实际情况处理注册失败的逻辑，例如显示错误消息
        }
    } catch (error) {
        // 处理错误
        console.error('注册失败:', error.message);
        showCustomToast(`注册失败 ${error.message}`);
        // 根据实际情况进行错误处理，例如显示错误消息
    }
}

// 获取推荐文章
// script.js

// Assuming you have a function to format the date
// Replace formatDate() with your actual function
// Assuming you have a function to format the date
// Replace formatDate() with your actual function
function formatDate(dateString) {
    const options = { year: 'numeric', month: '2-digit', day: '2-digit', hour: 'numeric', minute: 'numeric', second: 'numeric', timeZone: 'Asia/Shanghai' };
    return new Date(dateString).toLocaleDateString('zh-CN', options).replace(/\//g, '-');
}

// Assuming you have a function to fetch recommended articles
// Replace fetchRecommendedArticles() with your actual function
function fetchRecomArticles() {
    const apiUrl = 'http://8.130.162.251:3389/api/v1/article/recom';
    console.log(apiUrl)
    fetch(apiUrl)
        .then(response => {
            const contentType = response.headers.get('content-type');
            if (!contentType || !contentType.includes('application/json')) {
                throw new Error('服务器返回的不是有效的 JSON 数据');
            }

            return response.json();
        })
        .then(data => {
            const articleList = document.getElementById('articleList');
            console.log(data)
            data.data.forEach(article => {
                const listItem = document.createElement('li');
                const link = document.createElement('a');
                link.href = `article_content.html?id=${encodeURIComponent(article.ID)}`;
                link.textContent = article.title;

                const time = document.createElement('span');
                time.textContent = formatDate(article.CreatedAt);
                time.classList.add('time'); // Add the 'time' class to the time element
                listItem.appendChild(link);
                listItem.appendChild(time); // Append the time element
                articleList.appendChild(listItem);

                articleList.appendChild(listItem);
            });
        })

        .catch(error => {
            console.error('请求失败:', error.message);
            // 根据实际情况进行错误处理，例如显示错误消息
        });
}


// 获取并返回分类列表
async function fetchData() {
    try {
        const response = await fetch('http://8.130.162.251:3389/api/v1/categories?pagesize=10&pagenum=1');
        const data = await response.json();
        return data.data;
    } catch (error) {
        console.error('获取数据时发生错误:', error);
    }
}

// 更新下拉菜单的函数
async function updateDropdown() {
    const dropdown = document.getElementById('dynamicDropdown');
    dropdown.innerHTML = ''; // 清空现有内容

    const categories = await fetchData();

    categories.forEach(category => {
        const link = document.createElement('a');
        link.href = `cate_articles.html?cname=${encodeURIComponent(category.name)}&cid=${encodeURIComponent(category.id)}`;
        link.textContent = category.name;

        // 阻止默认链接行为，以便使用自定义链接
        link.addEventListener('click', (event) => {
            event.preventDefault();
            window.location.href = link.href; // 手动进行页面跳转
        });

        dropdown.appendChild(link);
    });
}

// 修改“个人中心”下三个超链接的链接
function changeLink() {
    const username = getUsername()
    // 找到具有特定 id 属性的 <a> 元素并修改其 href 属性
    const loginLink = document.getElementById('loginLink');
    const artisLink = document.getElementById('artisLink');
    const editLink = document.getElementById('editLink');
    
    if (username == "") { // 未登录 

        if (loginLink) {
            loginLink.href = 'login.html';
            loginLink.textContent = '登录';

        }
        if (artisLink) {
            artisLink.href = `javascript:showCustomToast('请在登录后使用本功能')`;

        }
        if (editLink) {
            editLink.href = `javascript:showCustomToast('请在登录后使用本功能')`;

        }
    }
    else { // 已登录
        if (loginLink) {
            loginLink.href = 'javascript:logout()';
            loginLink.textContent = '退出登录';
        }
        if (artisLink) {
            artisLink.href = `user_articles.html?username=${username}`;

        }
        if (editLink) {
            editLink.href = `edit.html`;

        }
    }



}


// // 修改个人中心下拉菜单的内容
// function setUserArtLink() {
//     const token = sessionStorage.getItem('token');


//     const loginLink = document.getElementById('loginLink');
//     const artisLink = document.getElementById('artisLink');
//     const editLink = document.getElementById('editLink');


//     if (token!=null) {
//         // 如果用户已登录，可以修改为实际的用户中心链接
//         loginLink.textContent = '退出登录';
//         sessionStorage.removeItem('token');
//         loginLink.href = 'javascript:logout();'; // 替换为退出登录的函数
//         artisLink.href = '#';  // 替换为实际的查看文章链接
//     } else {
//         // 如果用户未登录，保持链接不变
//         showCustomToast('您还没有登录哦! ');
//         loginLink.href = 'login.html';
//         artisLink.href = '#';  // 未登录时可以设为 # 或其他适当的操作

//     }
// }



// 从token中获取username
function getUsername() {
    // 获取 token
    const token = sessionStorage.getItem('token');
    // 如果token为空说明未登录
    if (token == null) {
        return ""
    }
    // 解码 token 获取用户名信息
    const decodedToken = atob(token.split('.')[1]);
    const tokenData = JSON.parse(decodedToken);
    return tokenData.username;

}



// 提交编辑的文章
function submitArticle(event) {
    event.preventDefault(); // 阻止表单的默认提交行为

    // 获取 token
    const token = sessionStorage.getItem('token');


    const username = getUsername()
    // 从8.130.162.251:3389/api/v1/user/id/username获取uid
    fetch(`http://8.130.162.251:3389/api/v1/user/id/${username}`)
        .then(response => response.json())
        .then(user => {
            const uid = user.uid;

            const selectedCategoryId = document.getElementById("articleCategory").value;
            const cid = parseInt(selectedCategoryId, 10); // 把字符串类型转化为uint类型

            const articleData = {
                title: document.getElementById("articleTitle").value,
                desc: document.getElementById("articleIntroduction").value,
                content: document.getElementById("articleContent").value,
                img: "",
                uid: uid, // Use the obtained uid here
                cid: cid
            };

            // Submit article data
            fetch("http://8.130.162.251:3389/api/v1/article/add", {
                method: "POST",
                headers: {
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify(articleData)
            })
                .then(response => response.json())
                .then(data => {
                    console.log("提交成功", data);
                    alert("提交成功！");
                })
                .catch(error => {
                    console.error("提交失败", error);
                });
        })
        .catch(error => {
            console.error("获取uid失败", error);
        });
}

// 填充分类下拉框
function populateDropdown(categoryDict) {
    var dropdown = document.getElementById("articleCategory");

    // 清空下拉框，以防已有的选项
    dropdown.innerHTML = "";

    // 遍历字典中的每个键值对
    for (var category in categoryDict) {
        if (categoryDict.hasOwnProperty(category)) {
            // 创建一个新的选项元素
            var option = document.createElement("option");

            // 设置选项的值和显示文本
            option.value = categoryDict[category];
            option.text = category;

            // 将选项添加到下拉框
            dropdown.add(option);
        }
    }
}




// user_articles.html相关

function getUsernameFromUrl() {


    const searchParams = new URLSearchParams(window.location.search);
    const username = searchParams.get('username');
    console.log(username)
    // 将username写入标题
    document.getElementById('pageTitle').innerText = username;

    return username;
}





// 获取某个用户下的文章列表
function fetchUserArticles(username) {
    // 构造 GET 请求的地址

    // 获取用户id
    fetch(`http://8.130.162.251:3389/api/v1/user/id/${username}`)
        .then(response => response.json())
        .then(user => {
            const uid = user.uid;
            // 用于获取用户下文章列表api
            const apiUrl = `http://8.130.162.251:3389/api/v1/article/ulist/${uid}?pagesize=10&pagenum=1`;
            console.log(apiUrl)
            // 发送 GET 请求
            // fetch(apiUrl)
            fetch(apiUrl, {
                method: 'GET',
                headers: {
                    // 添加 CORS 相关 headers
                    // 暂时好像都不需要
                    // 'Origin': 'http://localhost:3389'
                    // 其他可能需要的 headers...
                },
            })

                .then(response => {
                    // 检查响应是否为 JSON 格式
                    const contentType = response.headers.get('content-type');
                    if (!contentType || !contentType.includes('application/json')) {
                        throw new Error('服务器返回的不是有效的 JSON 数据');
                    }

                    return response.json();

                })
                .then(data => {
                    // 处理接口返回的数据
                    console.log('接口返回的数据:', data);

                    // 提取每篇文章的标题和id
                    const articles = data.data.map(article => ({
                        title: article.title,
                        id: article.ID,
                    }));

                    // 在这里添加处理文章标题和id的逻辑，例如展示在页面上
                    console.log('文章列表:', articles);

                    const articleList = document.getElementById('articleList');
                    articles.forEach(article => {
                        const listItem = document.createElement('li');
                        const link = document.createElement('a');
                        link.href = `article_content.html?title=${encodeURIComponent(article.title)}&id=${encodeURIComponent(article.id)}`;
                        link.textContent = article.title;
                        listItem.appendChild(link);
                        articleList.appendChild(listItem);
                    });

                })
                .catch(error => {
                    // 处理错误
                    console.error('请求失败:', error.message);

                    // 根据实际情况进行错误处理，例如显示错误消息
                });
        })
        .catch(error => {
            console.error("获取uid失败", error);
        });

}




// 用于提示的函数，使用前需要在页面内添加<div id="customToast" class="toast"></div>
function showCustomToast(message) {
    var toast = document.getElementById('customToast');
    toast.innerText = message;
    toast.style.animation = 'slideIn 0.5s ease-out'; // 添加动画效果
    toast.style.display = 'block';

    setTimeout(function () {
        toast.style.animation = 'slideOut 0.5s ease-in'; // 添加退出动画效果
        setTimeout(function () {
            toast.style.display = 'none';
            toast.style.animation = ''; // 重置动画效果
        }, 500);
    }, 1000);
}
