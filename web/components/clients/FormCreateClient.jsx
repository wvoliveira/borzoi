import React from "react";
import {TextField} from "@mui/material";
import {LoadingButton} from "@mui/lab";
import {ClientAPI} from "../../lib/api/client";
import useAuth from "../../lib/utils/useAuth";

export default function FormCreateClient() {
    const [loading, setLoading] = React.useState(false);
    const [errors, setErrors] = React.useState([]);

    const [name, setName] = React.useState("");
    const [description, setDescription] = React.useState("");

    const handleName = React.useCallback((e) => setName(e.target.value), []);
    const handleDescription = React.useCallback((e) => setDescription(e.target.value), []);

    const {mutate: mutateFindAll} = ClientAPI.FindAll();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

        console.log("name: " + name);
        console.log("description: " + description);

        try {
            const {data, status} = await ClientAPI.Create(name, description);
            if (status !== 200 && status !== 500) {
                setErrors(data.message);
                console.log(data.message);
            }
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
            setName("")
            setDescription("")
            await mutateFindAll()
        }
    };

    return (
        <form>
            <TextField
                name="name"
                label="Name"
                type="text"
                placeholder="Name"
                value={name}
                onChange={handleName}
                size="small"
                required
            />
            {" "}
            <TextField
                name="description"
                label="Description"
                type="text"
                placeholder="Description"
                value={description}
                onChange={handleDescription}
                size="small"
                required
            />
            {" "}
            <LoadingButton type="submit" onClick={handleSubmit} variant="contained" loading={loading}>
                Create
            </LoadingButton>
            {" "}
            {errors ? errors : ""}
        </form>
    );
}
