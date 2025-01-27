import SignUpForm from "./components/signUpForm";

export default function Page() {
    return (
        <main>
            <h1>Sign Up</h1>
            <SignUpForm/>
            <a href="/signin">Sign In Instead</a>
        </main>
    )
}