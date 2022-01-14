const SuperSecretPassword = "TrustNo1";  // you can't see it

window.onload = function () {
    const levers = Array.from(document.querySelectorAll('.levers input'));
    const boxes = Array.from(document.querySelectorAll('.check-buttons input'));
    const controls = [boxes, levers].flat();

    changeControlAvailability('disable', controls);

    const launchButton = document.querySelector('.launch-button');
    launchButton.disabled = true;
    const passwordInput = document.querySelector('.password-input');
    const passwordButton = document.querySelector('.ok-button');

    passwordButton.addEventListener('click', function () {
        let value = passwordInput.value;
        passwordInput.value = '';
        if(value === SuperSecretPassword) {
            console.log('Password ok.');
            changeControlAvailability('enable', controls);
            passwordButton.disabled = true;
            passwordInput.disabled = true;
        } else {
            console.log('Wrong password!');
        }
    });

    function inputCheck() {
        for(const i in boxes) {
            if(!boxes[i].checked) {
                return false;
            }
        }
        for(const i in levers) {
            if(levers[i].value !== '100') {
                return false;
            }
        }

        return true;
    }

    function inputEvent() {
        launchButton.disabled = !inputCheck();
    }

    for(const i in controls) {
        controls[i].onchange = inputEvent;
    }

    launchButton.addEventListener('click', animation);
};

function changeControlAvailability(action, controls) {
    action = action === "enable";

    for(let i in controls) {
        controls[i].disabled = !action;
    }
}

function animation() {
    let ship = document.querySelector('.rocket');

    ship.animate([
        {transform: "rotate(30deg)"},
        {transform: "rotate(30deg) translate(0, -100vh)"}
    ], {
        duration: 2000
    });
    ship.style.transform = "rotate(30deg) translate(0, -100vh)";
}
