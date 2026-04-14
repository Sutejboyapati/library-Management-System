describe('Library login and browse flow', () => {
  it('loads demo data, logs in, and opens the catalog', () => {
    cy.visit('/login');
    cy.get('[data-testid="seed-demo-data"]').click();
    cy.get('[data-testid="login-username"]').type('admin');
    cy.get('[data-testid="login-password"]').type('admin123');
    cy.get('[data-testid="login-submit"]').click();
    cy.url().should('include', '/main');
    cy.visit('/main/books');
    cy.get('[data-testid="book-search-input"]').type('clean');
    cy.get('[data-testid="book-search-button"]').click();
    cy.get('[data-testid="book-card"]').should('exist');
  });
});
