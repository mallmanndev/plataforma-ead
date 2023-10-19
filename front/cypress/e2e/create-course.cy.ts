Cypress.config()

describe("Test Create Course", () => {
    it("When user type is not admin, redirect to login")

    it("When user type is admin, load page successfully", () => {
        cy.login("matheusvmallmann@gmail.com", "12345678")
        cy.visit('/create-course')
        cy.contains("Criar Curso")
    })

    it("When course title and description is invalid, show error message", () => {
        cy.login("matheusvmallmann@gmail.com", "12345678")
        cy.visit('/create-course')

        cy.get("input[name=name").type("Abc")
        cy.get("textarea[name=description]").type("Test")
        cy.get("button[type=submit]").click()
    })

    it("When course is valid, create course successfully")
    // cy.visit('http://localhost:3002/create-course')


})