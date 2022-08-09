let formSignup = document.querySelector("#form_signup")
let formLogin = document.querySelector("#form_login")
localStorage.clear()

formSignup.addEventListener('submit', async (e) => {
    e.preventDefault()
    try {
        let response = await fetch(`http://localhost:8080/auth/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: document.querySelector("#name_signup").value,
                email: document.querySelector("#email_signup").value,
                password: document.querySelector("#pass_signup").value
            })
        })
        console.log(localStorage.getItem('token'))
        let result = await response.json()
        console.log(result)
        if (result.status == true) {
            localStorage.setItem('token', result.data.token)
            window.location = './index.html'
        }
        
    } catch (e) {
        console.log(err)
    }
})

formLogin.addEventListener('submit', async (e) => {
    e.preventDefault()
    try {
        let response = await fetch(`http://localhost:8080/auth/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: document.querySelector("#email_login").value,
                password: document.querySelector("#pass_login").value,
            })
        })
        let result = await response.json()
        console.log(result)
        if (result.status == true) {
            localStorage.setItem('token', result.data.token)
            window.location = './index.html'
        }
        
    } catch (e) {
        console.log(err)
    }
})

