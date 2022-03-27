import * as cookie from "../utils/cookie";

export function IsAuthenticated() {
    if (typeof document !== "undefined") {
        let session = cookie.GetCookie("session")
        if (session === null || session === "") {
            return false;
        }
    }
    return true;
}
