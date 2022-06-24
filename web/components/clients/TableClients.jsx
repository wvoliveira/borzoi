import React from "react";
import { ClientAPI } from "../../lib/api/client";
import { DataGrid } from "@mui/x-data-grid";
import Rating from "@mui/material/Rating";
import { format } from "date-fns";
import Button from '@mui/material/Button';
import Skeleton from '@mui/material/Skeleton';
import Stack from '@mui/material/Stack';

export default function TableClients() {
    const columns = [
        { field: "id", headerName: "ID", width: 50 },
        { field: "name", headerName: "Name", flex: 1 },
        { field: "description", headerName: "Description", flex: 1.5 },
        {
            field: "score",
            headerName: "Score",
            flex: 1,
            renderCell: (params) => (
                <div>
                    <Rating value={params.value} readOnly />
                </div>
            ),
        },
        {
            field: "created_at",
            headerName: "Created at",
            flex: 1,
            renderCell: (params) => (
                <>{format(Date.parse(params.value), "dd/MM/yyyy HH:mm:ss")}</>
            ),
        },
        {
            field: "updated_at",
            headerName: "Updated at",
            flex: 1,
            renderCell: (params) => (
                <>{format(Date.parse(params.value), "dd/MM/yyyy HH:mm:ss")}</>
            ),
        },
    ];

    const [page, setPage] = React.useState(0);
    const [limit, setLimit] = React.useState(5);
    const [total, setTotal] = React.useState(0);
    const [rows, setRows] = React.useState([]);

    const [loading, setLoading] = React.useState(true);
    const [selectionModel, setSelectionModel] = React.useState([]);
    const prevSelectionModel = React.useRef(selectionModel);

    const getClients = async () => {
        const result = await ClientAPI.FindAll(page, limit);
        const body = await result.data;
        const data = body.data;

        setRows(data.clients);
        setTotal(data.total);
    };

    const handleUpdateAllRows = () => {
        getClients();
    };

    React.useEffect(() => {
        setLoading(true);
        getClients();
        setLoading(false)

        setTimeout(() => {
            setSelectionModel(prevSelectionModel.current);
        });
    }, [loading, page, limit]);

    return (
        <div>
        <button size="small" onClick={handleUpdateAllRows}>
          Refresh
        </button>
        <div>
            {loading ? (<>
                <Stack>
                    <Skeleton variant="rectangular" height={50} />
                    {Array.from({length: 10}, (_, i) => (
                        <Skeleton key={i} variant="text" height={50} />
                    ))}
                </Stack>
            </>) : <>
            <table>
            </table>
            </>}
        </div>
        </div>
    );
}
