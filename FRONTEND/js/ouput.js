if (localStorage.getItem('token') == null) {
    window.location = './login.html'
}

let output = document.querySelector('#shortenurl')
let longBox = document.querySelector('#longBox')

output.value = `${localStorage.getItem('shortUrl')}`
longBox.innerHTML = `Long URL: <a href='${localStorage.getItem('longUrl')}' target="_blank">${localStorage.getItem('longUrl')}</a>`

let copyBtn = document.querySelector('.copy')
copyBtn.addEventListener('click', (e) => {
    let copyTextarea = document.querySelector('#shortenurl');
    copyTextarea.focus();
    copyTextarea.select();
    try {
        var successful = document.execCommand('copy');
        var msg = successful ? 'successful' : 'unsuccessful';
        console.log('Copying text command was ' + msg);
    } catch (err) {
        console.log('Oops, unable to copy');
    }
})
