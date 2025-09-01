document.addEventListener('DOMContentLoaded', function() {  
    // Existing form validation code...  
    const forms = document.querySelectorAll('form');  
    forms.forEach(form => {  
        form.addEventListener('submit', function(event) {  
            if (!form.checkValidity()) {  
                event.preventDefault();  
                event.stopPropagation();  
            }  
            form.classList.add('was-validated');  
        });  
    });  
  
    // Live search functionality  
    const searchInputs = document.querySelectorAll('input[name="q"]');  
    searchInputs.forEach(input => {  
        let searchTimeout;  
          
        input.addEventListener('input', function() {  
            clearTimeout(searchTimeout);  
            const query = this.value.trim();  
              
            // Debounce search requests  
            searchTimeout = setTimeout(() => {  
                if (query.length >= 2) {  
                    performLiveSearch(query);  
                }  
            }, 300);  
        });  
    });  
  
    // Keyboard shortcut for search (Ctrl+K or Cmd+K)  
    document.addEventListener('keydown', function(event) {  
        if ((event.ctrlKey || event.metaKey) && event.key === 'k') {  
            event.preventDefault();  
            const searchInput = document.querySelector('input[name="q"]');  
            if (searchInput) {  
                searchInput.focus();  
                searchInput.select();  
            }  
        }  
    });  
});  
  
function performLiveSearch(query) {  
    fetch(`/api/search?q=${encodeURIComponent(query)}`)  
        .then(response => response.json())  
        .then(data => {  
            console.log(`Found ${data.length} users matching "${query}"`);  
            // You could add a dropdown with search suggestions here  
        })  
        .catch(error => {  
            console.error('Search error:', error);  
        });  
}  
  
// Add search suggestions dropdown (optional enhancement)  
function showSearchSuggestions(users, inputElement) {  
    // Remove existing dropdown  
    const existingDropdown = document.querySelector('.search-suggestions');  
    if (existingDropdown) {  
        existingDropdown.remove();  
    }  
  
    if (users.length === 0) return;  
  
    // Create dropdown  
    const dropdown = document.createElement('div');  
    dropdown.className = 'search-suggestions';  
    dropdown.style.cssText = `  
        position: absolute;  
        background: white;  
        border: 1px solid #ddd;  
        border-radius: 4px;  
        box-shadow: 0 2px 8px rgba(0,0,0,0.1);  
        max-height: 200px;  
        overflow-y: auto;  
        z-index: 1000;  
        width: 100%;  
    `;  
  
    users.slice(0, 5).forEach(user => {  
        const item = document.createElement('div');  
        item.className = 'search-suggestion-item';  
        item.style.cssText = `  
            padding: 8px 12px;  
            cursor: pointer;  
            border-bottom: 1px solid #eee;  
        `;  
        item.innerHTML = `<strong>${user.name}</strong><br><small>${user.email}</small>`;  
          
        item.addEventListener('click', () => {  
            window.location.href = `/edit/${user._id}`;  
        });  
  
        dropdown.appendChild(item);  
    });  
  
    // Position dropdown  
    const rect = inputElement.getBoundingClientRect();  
    dropdown.style.top = (rect.bottom + window.scrollY) + 'px';  
    dropdown.style.left = rect.left + 'px';  
    dropdown.style.width = rect.width + 'px';  
  
    document.body.appendChild(dropdown);  
  
    // Close dropdown on outside click  
    document.addEventListener('click', function closeDropdown(event) {  
        if (!dropdown.contains(event.target) && event.target !== inputElement) {  
            dropdown.remove();  
            document.removeEventListener('click', closeDropdown);  
        }  
    });  
}  
