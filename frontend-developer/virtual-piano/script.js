let supported_keys = [
    // white keys
    "A",
    "S",
    "D",
    "F",
    "G",
    "H",
    "J",
    // black keys
    "W",
    "E",
    "T",
    "Y",
    "U",
];

function animateKey(id) {
    let keyItem = document.getElementById(id);
    keyItem.animate([
        { backgroundColor: "red" },
        { backgroundColor: keyItem.style.backgroundColor },
    ], {
        duration: 500,
    });

}

document.addEventListener("keydown", function (e) {
    let key = e.key.toUpperCase();
    if(supported_keys.includes(key)) {
        console.log("The '" + key + "' key is pressed.");

        // start the animation
        animateKey(e.code);

        // play the sound
        let src = "sounds/" + key + ".mp3";
        let sound = new Audio(src);
        sound.play();
    }
});