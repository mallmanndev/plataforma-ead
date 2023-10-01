import {faker} from "@faker-js/faker";

describe("Test Login", () => {
    it("Should show errors when inputs is invalid", () => {
        cy.visit('http://localhost:3002/login')

        cy.get('input[name=email]').type("matheus@email")
        cy.get('input[name=password]').type("12345")
        cy.get('button[type=submit]').click()

        cy.get('p[id=email-message]').should('be.visible')
        cy.get('p[id=email-message]').contains("Email inválido!")

        cy.get('p[id=password-message]').should('be.visible')
        cy.get('p[id=password-message]').contains("A senha deve conter mais de 8 digitos!")
    })

    it("Should show error when email is invalid", () => {
        cy.visit('http://localhost:3002/login')

        cy.get('input[name=email]').type(faker.internet.email())
        cy.get('input[name=password]').type(faker.internet.password())
        cy.get('button[type=submit]').click()

        cy.get('p[id=email-message]').should('not.exist')
        cy.get('p[id=password-message]').should('not.exist')

        cy.get('div[id=error-alert]').contains('User not found!')
    })

    it("Should show error when password is invalid", () => {
        cy.visit('http://localhost:3002/login')

        cy.get('input[name=email]').type("matheus@email.com")
        cy.get('input[name=password]').type(faker.internet.password())
        cy.get('button[type=submit]').click()

        cy.get('p[id=email-message]').should('not.exist')
        cy.get('p[id=password-message]').should('not.exist')

        cy.get('div[id=error-alert]').contains('Invalid password!')
    })

    it("Should show error when password is invalid", () => {
        cy.visit('http://localhost:3002/login')

        cy.get('input[name=email]').type("matheus@email.com")
        cy.get('input[name=password]').type("12345678")
        cy.get('button[type=submit]').click()

        cy.get('p[id=email-message]').should('not.exist')
        cy.get('p[id=password-message]').should('not.exist')

        cy.url().should('eq', 'http://localhost:3002/home')
    })
})