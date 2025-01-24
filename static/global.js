function createCookie(name,value,days) {
	if (days) {
		var date = new Date();
		date.setTime(date.getTime()+(days*24*60*60*1000));
		var expires = "; expires="+date.toGMTString();
	}
	else var expires = "";
	document.cookie = name+"="+value+expires+"; path=/";
}

function readCookie(name) {
	var nameEQ = name + "=";
	var ca = document.cookie.split(';');
	for(var i=0;i < ca.length;i++) {
		var c = ca[i];
		while (c.charAt(0)==' ') c = c.substring(1,c.length);
		if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
	}
	return null;
}

function eraseCookie(name) {
	createCookie(name,"",-1);
}

async function apiPost(endpoint, body) {
    const raw = JSON.stringify(body);
  
    const requestOptions = {
      method: "POST",
      body: raw,
      redirect: "follow"
    };
  
    res = await fetch("http://localhost:8080" + endpoint, requestOptions)
    return res
}

async function testSession() {
    sid = readCookie("session")
    if (sid == null) {
        return null
    }
    res = await apiPost("/api/auth/validate", {"sid": sid})
    if (res.status != 200) {
        eraseCookie("session")
        return null
    }
    return await res.json()
}

function signInSignOut() {
    eraseCookie("session")
    window.location.href = "/login"
}

window.addEventListener('load', async function() {
    let u = await testSession()
    if (this.location.pathname == "/login" && u) {
        this.location.href = "/"
    } else if (this.location.pathname != "/login" && !u) {
        this.location.href = "/login";
    }

    const el_SignInSignOut = this.document.getElementById("signin_signout")
    const el_UsernameText = this.document.getElementById("username_text")
    if (u) {
        el_SignInSignOut.innerHTML = "Sign Out"
        el_UsernameText.innerHTML = u.Username
    } 
    
    if (typeof onTemplateLoad !== "undefined") {
        onTemplateLoad(u)
    }
})