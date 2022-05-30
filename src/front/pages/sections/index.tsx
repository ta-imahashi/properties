import type { NextPage } from "next";
import { ApolloClient, gql, InMemoryCache } from "@apollo/client";
import { useEffect, useState } from "react";
import { CreateSection } from "components/CreateSection";
import { DeleteSection } from "components/DeleteSection";
import { DataGrid, GridColDef, GridValueGetterParams } from "@mui/x-data-grid";
import { format } from "date-fns";
import ja from "date-fns/locale/ja";
import { SectionTypeEnum } from "enums/SectionType";

const columns: GridColDef[] = [
  { field: "id", headerName: "ID", width: 50 },
  {
    field: "name",
    headerName: "区画名",
    width: 200,
    editable: true,
  },
  {
    field: "type",
    headerName: "区分",
    width: 70,
    editable: true,
    type: "singleSelect",
    valueOptions: Object.keys(SectionTypeEnum),
    valueFormatter: (params) => {
      const key: keyof typeof SectionTypeEnum = params.value;
      return SectionTypeEnum[key];
    },
  },
  {
    field: "property",
    headerName: "物件名",
    width: 200,
    editable: false,
    valueFormatter: (params) => params.value.name,
  },
  {
    field: "created_at",
    headerName: "作成日",
    sortable: false,
    width: 250,
    valueFormatter: (params) => {
      if (params.value) {
        const ts = Date.parse(params.value);
        const dt: Date = new Date(ts);
        return format(dt, "yyyy-MM-dd (EEEE) HH:mm:ss", { locale: ja });
      }
      return "";
    },
  },
  {
    field: "updated_at",
    headerName: "最終更新日",
    sortable: false,
    width: 250,
    valueFormatter: (params) => {
      if (params.value) {
        const ts = Date.parse(params.value);
        const dt: Date = new Date(ts);
        return format(dt, "yyyy-MM-dd (EEEE) HH:mm:ss", { locale: ja });
      }
      return "";
    },
  },
];

const Sections: NextPage = () => {
  const [rows, setRowData] = useState([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    const client = new ApolloClient({
      uri: "http://localhost:3000/graphql",
      cache: new InMemoryCache(),
    });
    client
      .query({
        query: gql`
          {
            getSection {
              id
              name
              type
              property_id
              property {
                name
              }
              created_at
              updated_at
            }
          }
        `,
      })
      .then((res) => {
        setRowData(res.data.getSection);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }, []);

  return (
    <div style={{ height: "88vh", width: "100%" }}>
      <CreateSection></CreateSection>
      <DeleteSection></DeleteSection>
      <DataGrid
        rows={rows}
        columns={columns}
        loading={isLoading}
        pagination
        pageSize={100}
        rowsPerPageOptions={[25, 50, 100, 250, 500, 1000]}
        // paginationMode='server'
        // checkboxSelection
        disableSelectionOnClick
      />
      {/* <DataTable columns={columns} rows={rows}></DataTable> */}
    </div>
  );
};

export default Sections;
