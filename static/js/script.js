
document.addEventListener('DOMContentLoaded', () => {
    const requestType = document.getElementById('request-type');
    const jsonFieldGroup = document.getElementById('json-field-group');

    requestType.addEventListener('change', () => {
        if (requestType.value === 'POST') {
            jsonFieldGroup.classList.remove('hidden');
        } else {
            jsonFieldGroup.classList.add('hidden');
        }
    });
});


document.getElementById("metrics-form").addEventListener("submit", async function (event) {
    event.preventDefault();

    // Собираем данные формы
    const formData = {
        request_url: document.getElementById("request-url").value,
        request_count: parseInt(document.getElementById("request-count").value),
        parallel: document.getElementById("parallel-checkbox").checked,
        request_type: document.getElementById("request-type").value,
        json_data: document.getElementById("json-input").value,
        randomize: document.getElementById("randomize-checkbox").checked
    };

    // Отправляем запрос на сервер
    const response = await fetch("/get-data", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    });

    // Обрабатываем ответ сервера
    const result = await response.json();
    console.log(result);
    alert(result.message || "Данные отправлены успешно!");
});



document.getElementById('request-url').addEventListener('focus', function() {
    if (!this.value) {
        this.value = 'http://';
    }
});

document.getElementById('json-input').addEventListener('input', function(event) {
    const text = this.value;
    const cursorPosition = this.selectionStart;

    if (text[cursorPosition - 1] === '{') {
        this.value = text.slice(0, cursorPosition) + '}' + text.slice(cursorPosition);
        this.setSelectionRange(cursorPosition, cursorPosition); 
    }
});