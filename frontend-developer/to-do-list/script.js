let taskListArray = storageLoadTasks();
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
    for(let task in taskListArray) {
        task = taskListArray[task];
        addTask(task.text, task.checked, false);
    }
}

function addTask(text, done=false, save=true) {
    let task = document.createElement('li');

    let span = document.createElement('span');
    span.innerText = text;
    span.className = "task";

    let chBox = document.createElement('input');
    chBox.type = "checkbox";
    chBox.addEventListener('input', function (e) {
        console.log(task);
        console.log(elementIndex(task));
        taskListArray[elementIndex(task)].checked = e.target.checked;
        storageSaveTasks();

        if(e.target.checked) {
            span.classList.add("task-done");
        } else {
            span.classList.remove("task-done");
        }
    });

    if(done) {
        span.classList.add("task-done");
        chBox.checked = true;
    }

    let delBtn = document.createElement('button');
    delBtn.className = "delete-btn";
    delBtn.addEventListener("click", function () {
        taskListArray.splice(elementIndex(task), 1);
        storageSaveTasks();

        task.remove();
    });

    task.appendChild(chBox);
    task.appendChild(span);
    task.appendChild(delBtn);

    taskList.appendChild(task);

    if(save) {
        taskListArray.push({text: text, done: done});
        storageSaveTasks();
    }
}

function elementIndex(element) {
    let index = -1;

    while(element) {
        if(element.nodeName.toUpperCase() === 'LI') {
            ++index;
        }
        element = element.previousSibling;
    }

    return index;
}

function storageSaveTasks() {
    localStorage.setItem("tasks", JSON.stringify(taskListArray));
}

function storageLoadTasks() {
    return JSON.parse(localStorage.getItem("tasks")) || [];
}

function storageClearTasks() {
    localStorage.removeItem("tasks");
}
