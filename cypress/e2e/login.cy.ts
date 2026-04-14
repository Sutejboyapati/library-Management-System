describe('Login page', () => {
  it('fills login form and navigates to dashboard', () => {
    cy.intercept('POST', '/api/login', {
      statusCode: 200,
      body: {
        message: 'Login successful',
        token: 'x.eyJ1c2VySWQiOjEsInJvbGUiOiJ1c2VyIn0.y',
        username: 'student1',
      },
    }).as('loginRequest');

    cy.visit('/login');
    cy.get('[data-cy="username-input"]').type('student1', { force: true });
    cy.get('[data-cy="password-input"]').type('password123', { force: true });
    cy.get('[data-cy="login-submit"]').click({ force: true });

    cy.wait('@loginRequest');
    cy.url().should('include', '/main/dashboard');
  });

  it('shows validation state for empty submit', () => {
    cy.visit('/login');
    cy.get('[data-cy="login-submit"]').click();
    cy.contains('Please enter username and password.').should('be.visible');
  });
});
