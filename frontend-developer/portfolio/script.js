function popupShow(title, text) {
    let popup = document.querySelector('#popup');
    let popupWindow = document.querySelector('#popup .window');
    popupWindow.children[0].innerHTML = title;
    popupWindow.children[1].innerHTML = text;

    popup.style.display = "block";
    popupWindow.animate([
        {top: "-50%"},
        {top: "60%", offset: 0.8, easing: "ease-in-out"},
        {top: "50%", easing: "ease-in-out"}
    ], {
        duration: 500
    });
    popup.animate([
        {opacity: 0},
        {opacity: 1, offset: 0.5},
        {opacity: 1}
    ], {
        duration: 500
    });
}

function popupHide() {
    let popup = document.querySelector('#popup');
    let popupWindow = document.querySelector('#popup .window');

    popupWindow.animate([
        {top: "50%"},
        {top: "60%", offset: 0.2, easing: "ease-in-out"},
        {top: "-50%", easing: "ease-in-out"}
    ], {
        duration: 500
    });
    popup.animate([
        {opacity: 1},
        {opacity: 1, offset: 0.5},
        {opacity: 0, display: "none"}
    ], {
        duration: 500
    }).onfinish = function () {
        popup.style.display = "none";
    };
}

window.onload = function () {
    document.querySelector('#popup button.close').addEventListener('click', popupHide);
    document.querySelector('#popup #popup-bg').addEventListener('click', popupHide);

    let btns = document.querySelectorAll('.open-window');
    let openWindowFunc = function () {
        let sbl = this.nextSibling;  // text
        sbl = sbl.nextSibling;  // div
        let title = sbl.children[0].innerHTML;
        let text = sbl.children[1].innerHTML;
        popupShow(title, text);
    }
    btns.forEach(function (btn) {
        btn.addEventListener('click', openWindowFunc);
    });

    // this is for showing hamburger menu
    let hamburgerIcon = document.querySelector('.hamburger-icon');
    let hamburgerMenu = document.querySelector('ul.hamburger');
    hamburgerIcon.addEventListener('click', function () {
        if(hamburgerIcon.classList.toggle('clicked')) {
            hamburgerMenu.style.display = 'inline-block';
            hamburgerMenu.animate([
                {opacity: 0},
                {opacity: 1}
            ], {
                duration: 500
            });
        } else {
            hamburgerMenu.animate([
                {opacity: 1},
                {opacity: 0}
            ], {
                duration: 500
            }).onfinish = function () {
                hamburgerMenu.style.display = 'none';
            };
        }
    });
}
