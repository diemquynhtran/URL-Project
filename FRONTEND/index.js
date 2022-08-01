
let inputValue = ''
let input = document.querySelector('#input')
let urlContainer = document.querySelector('#urlbox')

const handleChange = (e) => {
    const { name, value } = e.target
    //Detructuring
    inputValue = value
}
input.addEventListener('blur', handleChange)

let myForm = document.querySelector('#form')
myForm.addEventListener('submit', async (e) => {
    e.preventDefault()
    try {
        let res = await fetch(`http://localhost:8080/url/?long=${inputValue}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        let result = await res.json()
        console.log(result.data)
        localStorage.setItem('shortUrl',result.data.short)
        localStorage.setItem('longUrl',result.data.long)
        window.location = './output.html'
    } catch (err) {
        console.log(err)
    }
})


