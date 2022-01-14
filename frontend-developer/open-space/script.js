const SuperSecretPassword = "TrustNo1";  // you can't see it

window.onload = function () {
    const controls = [
        Array.from(document.querySelectorAll('.levers input')),
        Array.from(document.querySelectorAll('.check-buttons input')),
        document.querySelector('.launch-button')
    ].flat();

    changeControlAvailability('disable', controls);

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
};

function changeControlAvailability(action, controls) {
    action = action === "enable";

    for(let i in controls) {
        controls[i].disabled = !action;
    }
}
