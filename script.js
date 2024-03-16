document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('survey-form');
    form.onsubmit = function(e) {
        e.preventDefault(); // Prevent the default form submission

        // Gather form data
        const formData = {
            name: document.getElementById('name').value,
            email: document.getElementById('email').value,
            age: document.getElementById('age').value,
            role: document.getElementById('role').value,
            favoriteFeature: document.getElementById('favorite-feature').value,
            improvements: Array.from(document.querySelectorAll('input[name="improvements"]:checked')).map(el => el.value),
            comments: document.getElementById('comments').value
        };

        // Send the form data to the Go backend
        fetch('http://localhost:8080/submit-survey', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert('An error occurred: ' + data.error);
            } else {
                alert('Success: ' + data.message);
                form.reset(); // Reset the form after successful submission
            }
        })
        .catch(error => {
            alert('An error occurred: ' + error);
        });
    };
});
