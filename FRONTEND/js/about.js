if (localStorage.getItem('token') == null) {
    window.location = './login.html'
}
else {
    async function getData() {
        try {
            let res = await fetch(`http://localhost:8080/url/urls`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': localStorage.getItem('token')
                }
            })
            let result = await res.json()
            console.log(result)
            return result
        } catch (err) {
            console.log(err)
        }
    }

    const tableBody = document.getElementById('table-data')

    const renderRow = async () => {
        let data = await getData()
        let html = ''
        data.map((item, index) => {
            html += `
            <tr>
            <td>${index + 1}</td>
            <td>${item.long}</td>
            <td><a href=${item.short} target="blank">${item.short}</a></td>
            <td>${item.click}</td>
            <td>${item.user.name}</td>
            <td><button data-item="${item.snowflake}" class="btn-status">${item.status}</button></td>
            <td><button data-item="${item.snowflake}" class="btn-delete">Delete</button></td>
        </tr>
            `
        })
        tableBody.innerHTML = html
        let btnStatus = document.querySelectorAll(".btn-status")
        let btnDels = document.querySelectorAll(".btn-delete")
        btnStatus.forEach(item => item.addEventListener('click', handleChangeStatus))
        btnDels.forEach(item => item.addEventListener('click',handleDelete))
    }
    const handleChangeStatus = async (e) => {
        e.preventDefault()
        try {
            let id = e.target.attributes[0].value
            let res = await fetch(`http://localhost:8080/url/edit/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': localStorage.getItem('token')
                }
            })
            window.location='./about.html'
        } catch (err) {
            console.log(err)
        }
    }

    const handleDelete = async (e) => {
        e.preventDefault()
        try {
            let id = e.target.attributes[0].value
            let res = await fetch(`http://localhost:8080/url/delete/${id}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': localStorage.getItem('token')
                }
            })
            window.location='./about.html'
        } catch (err) {
            console.log(err)
        }
    }

renderRow()
}
