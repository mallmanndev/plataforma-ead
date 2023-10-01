describe('Test Home screen', () => {
    it('Should redirect to home when user is not logged', () => {
        cy.visit('http://localhost:3002/home')
        cy.url().should('eq', 'http://localhost:3002/login')
    })

    it('Should redirect to login when token is invalid', () => {
        cy.setCookie("token", "invalid-token")
        cy.visit('http://localhost:3002/home')
        cy.url().should('eq', 'http://localhost:3002/login')
    })

    it('Should load page when token is invalid', () => {
        cy.setCookie("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjkxMTFiZmZkLTczZDktNDlkOC1iMzJjLTQ4MzUzNjc0ZGMwNiIsIm5hbWUiOiJNYXRoZXVzIE1hbGxtYW5uIiwiZW1haWwiOiJtYXRoZXVzQGVtYWlsLmNvbSIsInBob25lIjoiNTU5OTkwNDgyMjMiLCJ0eXBlIjoic3R1ZGVudCIsImlhdCI6MTY5NjE5NDY4MiwiZXhwIjoxNjk2MjgxMDgyfQ.OIB64eNRwU-hQSOllkQwZfcyDGiXjRNLvVYvLb2UtQs")
        cy.visit('http://localhost:3002/home')

        cy.get('button[id=user-avatar]').click()

        cy.get('span[id=sigla-avatar]').contains('MM')
        cy.get('p[id=user-name]').contains('Matheus Mallmann')
        cy.get('p[id=user-email]').contains('matheus@email.com')
    })
})