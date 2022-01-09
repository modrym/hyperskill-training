let supported_keys = ["KeyA",
    "KeyS",
    "KeyD",
    "KeyF",
    "KeyG",
    "KeyH",
    "KeyJ",
];

document.addEventListener("keydown", function (e) {
    if(supported_keys.includes(e.code)) {
        console.log("The '" + e.key.toUpperCase() + "' key is pressed.");
    }
});