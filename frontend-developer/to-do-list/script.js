let taskList = document.getElementById('task-list');

document.getElementById('add-task-button').addEventListener('click', function () {
    let input = document.getElementById('input-task');

    let value = input.value.trim();

    if(value) {
        addTask(value);
    }

    input.value = '';
});

window.onload = function () {
    addTask("Email David");
    addTask("Create ideal");
    addTask("Set up A/B test");
}

function addTask(text) {
    let task = document.createElement('li');

    let span = document.createElement('span');
    span.innerText = text;
    span.className = "task";
    let chBox = document.createElement('input');
    chBox.type = "checkbox";

    let delBtn = document.createElement('button');
    delBtn.className = "delete-btn";
    delBtn.addEventListener("click", function () {
        task.remove();
    });

    task.appendChild(chBox);
    task.appendChild(span);
    task.appendChild(delBtn);

    taskList.appendChild(task);
}


