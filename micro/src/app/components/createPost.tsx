"use client";

import Cookies from "js-cookie";
import { FormEvent, useState } from "react";
import { CreatePost, postResponseToPost } from "../lib/api";
import { Post } from "../types";

export default function CreatePostBox({action}: {action: (p: Post) => void}) {
    const [status, setStatus] = useState<string>("")
    const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const form = event.currentTarget
        const formElements = form.elements as typeof form.elements & {
            textInput: {value: string}
        }
        const sess = Cookies.get("session")
        if (sess === undefined) {
            setStatus("log in first!")
            return;
        } else {
            CreatePost(sess, formElements.textInput.value).then(p => {
                if (!p) {
                    return;
                }
                postResponseToPost(p).then(po => {
                    action(po);
                    formElements.textInput.value = "";
                    setStatus("done!")
                })
            })
        }
    }

    return (
        <form onSubmit={handleSubmit}>
            <p>{status}</p>
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