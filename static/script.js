let input = document.querySelector('.goto-input');
input.addEventListener('keypress', e => {
    if (e.keyCode === 13) {
        let val = input.value;
        if (/^[0-9]+$/.test(val)) {
            window.location = '/' + val;
        } else {
            window.location = '/';
        }
    }
});