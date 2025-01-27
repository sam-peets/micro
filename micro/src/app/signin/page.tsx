import SignInForm from "./components/signInForm";

export default function Page() {
    return (
        <main>
            <h1>Sign In</h1>
            <SignInForm/>
            <a href="/signup">Sign Up Instead</a>
        </main>
    )
}