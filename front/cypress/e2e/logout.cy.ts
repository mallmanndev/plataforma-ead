describe('Test logout', () => {
    it('Should logout', () => {
        cy.setCookie("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjkxMTFiZmZkLTczZDktNDlkOC1iMzJjLTQ4MzUzNjc0ZGMwNiIsIm5hbWUiOiJNYXRoZXVzIE1hbGxtYW5uIiwiZW1haWwiOiJtYXRoZXVzQGVtYWlsLmNvbSIsInBob25lIjoiNTU5OTkwNDgyMjMiLCJ0eXBlIjoic3R1ZGVudCIsImlhdCI6MTY5NjE5NDY4MiwiZXhwIjoxNjk2MjgxMDgyfQ.OIB64eNRwU-hQSOllkQwZfcyDGiXjRNLvVYvLb2UtQs")
        cy.visit('http://localhost:3002/home')

        cy.get('button[id=user-avatar]').click()

        cy.get('div[id=logout]').click()
        cy.url().should('eq', 'http://localhost:3002/login')
    })
})