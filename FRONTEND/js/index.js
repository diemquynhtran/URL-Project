
let inputValue = ''
let input = document.querySelector('#input')
let urlContainer = document.querySelector('#urlbox')

const handleChange = (e) => {
    const { name, value } = e.target
    inputValue = value
}
input.addEventListener('blur', handleChange)
console.log(inputValue)

let myForm = document.querySelector('#form')
if (localStorage.getItem('token') == null) {
    window.location = './login.html'
}
else {
    myForm.addEventListener('submit', async (e) => {
        e.preventDefault()
        try {
            let res = await fetch(`http://localhost:8080/url/?long=${inputValue}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': localStorage.getItem('token')
                }
            })
            let result = await res.json()
            localStorage.setItem('shortUrl', result.data.short)
            localStorage.setItem('longUrl', result.data.long)
            window.location = './output.html'
        } catch (err) {
            console.log(err)
        }
    })

}
