export function IsAuthenticated() {
    if (typeof document !== "undefined") {
        let cookies = document.cookie;
        let session = cookies.split("=")[0];
        return session !== undefined;
    }
    return false;
}
