import {faker} from "@faker-js/faker";

describe('Test register', () => {
    it('Should return error when data is invalid', () => {
        cy.visit('http://localhost:3002/register')

        cy.get('input[name=name]').type("Ma")
        cy.get('input[name=phone]').type("555")
        cy.get('input[name=email]').type("matheus@email")
        cy.get('input[name=password]').type("12345")
        cy.get('input[name=cpassword]').type("54321")
        cy.get('button[type=submit]').click()

        cy.get('p[id=name-message]').should('be.visible')
        cy.get('p[id=name-message]').contains("Mínimo 5 caracteres.")

        cy.get('p[id=phone-message]').should('be.visible')
        cy.get('p[id=phone-message]').contains("Mínimo 9 caracteres.")

        cy.get('p[id=email-message]').should('be.visible')
        cy.get('p[id=email-message]').contains("Email inválido!")

        cy.get('p[id=password-message]').should('be.visible')
        cy.get('p[id=password-message]').contains("A senha deve conter mais de 8 digitos!")

        cy.get('p[id=cpassword-message]').should('be.visible')
        cy.get('p[id=cpassword-message]').contains("As senhas não conferem!")
    })

    it('Should show error alert when passwords not match', () => {
        cy.visit('http://localhost:3002/register')

        cy.get('input[name=name]').type("Matheus Mallmann")
        cy.get('input[name=phone]').type("(55) 99999-9999")
        cy.get('input[name=email]').type("matheus@email.com")
        cy.get('input[name=password]').type("12345678")
        cy.get('input[name=cpassword]').type("87654321")
        cy.get('button[type=submit]').click()

        cy.get('p[id=name-message]').should('not.exist')
        cy.get('p[id=email-message]').should('not.exist')
        cy.get('p[id=phone-message]').should('not.exist')
        cy.get('p[id=password-message]').should('not.exist')

        cy.get('p[id=cpassword-message]').should('be.visible')
        cy.get('p[id=cpassword-message]').contains("As senhas não conferem!")
    })

    it('Should show error alert when email is already registered', () => {
        cy.visit('http://localhost:3002/register')

        cy.get('input[name=name]').type("Matheus Mallmann")
        cy.get('input[name=phone]').type("(55) 99999-9999")
        cy.get('input[name=email]').type("matheus@email.com")
        cy.get('input[name=password]').type("12345678")
        cy.get('input[name=cpassword]').type("12345678")
        cy.get('button[type=submit]').click()

        cy.get('p[id=name-message]').should('not.exist')
        cy.get('p[id=email-message]').should('not.exist')
        cy.get('p[id=phone-message]').should('not.exist')
        cy.get('p[id=password-message]').should('not.exist')
        cy.get('p[id=cpassword-message]').should('not.exist')

        cy.get('div[id=error-alert]').contains('Email already registered!')
    })

    it('Should create user successfully', () => {
        cy.visit('http://localhost:3002/register')

        cy.get('input[name=name]').type(faker.person.fullName())
        cy.get('input[name=phone]').type("(55) 99999-9999")
        cy.get('input[name=email]').type(faker.internet.email())

        const pass = faker.internet.password()
        cy.get('input[name=password]').type(pass)
        cy.get('input[name=cpassword]').type(pass)
        cy.get('button[type=submit]').click()

        cy.get('p[id=name-message]').should('not.exist')
        cy.get('p[id=email-message]').should('not.exist')
        cy.get('p[id=phone-message]').should('not.exist')
        cy.get('p[id=password-message]').should('not.exist')
        cy.get('p[id=cpassword-message]').should('not.exist')

        cy.get('div[id=error-alert]').should('not.exist')
        cy.url().should('eq', 'http://localhost:3002/login')
    })
})