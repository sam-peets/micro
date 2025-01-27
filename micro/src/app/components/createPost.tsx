"use client";

import Cookies from "js-cookie";
import { FormEvent } from "react";
import { CreatePost } from "../lib/api";

export default function createPost() {
    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = event.currentTarget
        const formElements = form.elements as typeof form.elements & {
            textInput: {value: string}
        }
        const sess = Cookies.get("session")
        if (sess === undefined) {
            alert("log in first!")
            return;
        } else {
            const p = CreatePost(sess, formElements.textInput.value)
        }
    }
    return (
        <form onSubmit={handleSubmit}>
            <div style={{width: "100%", display: "flex"}}>
                <div style={{float: "left", width: "90%", height: "50px"}}>
                    <textarea id="textInput" name="text" placeholder="type" style={{width: "100%", height: "100%", resize: "none", boxSizing: "border-box"}}></textarea>
                </div>
                <div style={{float: "right", height: "50px", width: "10%"}}>
                    <button name="post" type="submit" style={{width: "100%", height: "100%"}}>Post</button>
                </div>
            </div>
        </form>
    )
}