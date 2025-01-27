import { useActionState } from "react"

function testAction(x: any) {
    console.log(x)
}

export default function createPost() {
    let handleSubmit = (event: any) => {
        event.preventDefault();
        console.log(event.target.elements.text.value)
    }
    return (
        <form onSubmit={handleSubmit}>
           <textarea name="text" placeholder="type" style={{width: "100%", resize: "none", }}></textarea>
           <button name="post" type="submit">Post</button>
        </form>
    )
}