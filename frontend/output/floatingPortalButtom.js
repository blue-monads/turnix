(function() {
    // Create button element
    const floatingButton = document.createElement('a');

    floatingButton.href = "/z/pages/portal";
    
    // Style the button
    Object.assign(floatingButton.style, {
        position: 'fixed',
        right: '20px',
        bottom: '20px',
        width: '2rem',
        height: '2rem',
        borderRadius: '50%',
        backgroundColor: '#ff9955',
        color: 'white',
        border: 'none',
        cursor: 'pointer',
        boxShadow: '0 2px 5px rgba(0,0,0,0.2)',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        zIndex: '9999',
        transition: 'transform 0.2s ease',
        padding: '0.4rem'
    });

    // Add hover effect
    floatingButton.onmouseover = () => {
        floatingButton.style.transform = 'scale(1.1)';
        floatingButton.style.backgroundColor = '#ff9900';
    };
    
    floatingButton.onmouseout = () => {
        floatingButton.style.transform = 'scale(1)';
        floatingButton.style.backgroundColor = '#ff9955';
    };

    const logoIcon = `
       <img id="floating-turnix-button" src="/z/lib/images/logo.png" alt="Back to turnix portal" />
    `;


    floatingButton.innerHTML = logoIcon;


    const updateSvgStyle = () => {
        const logo = floatingButton.querySelector('#floating-turnix-button');
        if (logo) {
            logo.style.width = '100%';
            logo.style.height = '100%';
            logo.style.display = 'block';
        }
    };


    // Add button to page when DOM is loaded
    document.addEventListener('DOMContentLoaded', () => {
        document.body.appendChild(floatingButton);
        updateSvgStyle(); // Style SVG after it's added to DOM
    });
})();