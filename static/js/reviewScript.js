document.querySelectorAll('.read-more').forEach(item => {
    item.addEventListener('click', event => {
      event.preventDefault();
      const content = item.previousElementSibling;
      content.textContent = {{.Content}};
    });
  });