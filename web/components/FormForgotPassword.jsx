import React from "react";
import {TextField} from "@mui/material";
import {LoadingButton} from "@mui/lab";

export default function FormForgotPassword() {
    const [loading, setLoading] = React.useState(false);
    const [error, setError] = React.useState("");
    const [sent, setSent] = React.useState(false);
    const [email, setEmail] = React.useState("");

    const handleEmailChange = React.useCallback(
        (e) => setEmail(e.target.value),
        []
    );

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

        try {
            // TODO: Send email.
            setSent(true)
        } catch (error) {
            console.error(error);
            setError(error)
        } finally {
            setLoading(false);
        }
    };

    return (
        <form>
            <TextField
                id="email"
                label="Email"
                type="email"
                placeholder="Email"
                value={email}
                onChange={handleEmailChange}
                size="small"
                required
            />
            {" "}
            <LoadingButton type="submit" onClick={handleSubmit} variant="contained" loading={loading}>
                Send
            </LoadingButton>
            {" "}
            {error ? error : ""}
            {sent ? "sent" : ""}
        </form>
    );
}
