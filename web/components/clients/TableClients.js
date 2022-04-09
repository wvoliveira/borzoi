import React from "react";
import { ClientAPI } from "../../lib/api/client";
import { DataGrid } from "@mui/x-data-grid";
import Rating from "@mui/material/Rating";
import { format } from "date-fns";

export default function TableClients() {
    const columns = [
        { field: "id", headerName: "ID", minWidth   : 30 },
        { field: "name", headerName: "Name", minWidth: 200 },
        { field: "description", headerName: "Description", minWidth: 400 },
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
    const [limit, setLimit] = React.useState(10);
    const [total, setTotal] = React.useState(0);
    const [rows, setRows] = React.useState([]);

    const [loading, setLoading] = React.useState(false);
    const [selectionModel, setSelectionModel] = React.useState([]);
    const prevSelectionModel = React.useRef(selectionModel);

    React.useEffect(() => {
        setLoading(true);

        const getClients = async () => {
            const result = await ClientAPI.FindAll(page, limit);
            const body = await result.data;
            const data = body.data;

            setRows(data.clients);
            setTotal(data.total);
        };

        getClients();
        setLoading(false);

        setTimeout(() => {
            setSelectionModel(prevSelectionModel.current);
        });
    }, [page, limit]);

    return (
        <div style={{ height: 500, width: "100%" }}>
            <DataGrid
                rows={rows}
                columns={columns}
                pagination
                pageSize={limit}
                rowsPerPageOptions={[5, 10, 20]}
                rowCount={total}
                paginationMode="server"
                onPageChange={(newPage) => {
                    prevSelectionModel.current = selectionModel;
                    setPage(newPage);
                }}
                onSelectionModelChange={(newSelectionModel) => {
                    setSelectionModel(newSelectionModel);
                }}
                onPageSizeChange={(newPageSize) => {
                    setLimit(newPageSize);
                }}
                selectionModel={selectionModel}
                loading={loading}
            />
        </div>
    );
}
