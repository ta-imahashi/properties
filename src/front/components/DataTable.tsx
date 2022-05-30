// import { Column, DataTable } from "components/DataTable";
// export type Section = {
//   id: number;
//   name: string;
//   created_at: string;
//   updated_at: string;
// };

// const columns: Column[] = [
//   {
//     value: "id",
//     label: "ID",
//     sortable: false,
//     align: "left",
//     filterable: true,
//   },
//   {
//     value: "name",
//     label: "区画名",
//     sortable: false,
//     align: "left",
//     filterable: true,
//   },
//   {
//     value: "created_at",
//     label: "作成日",
//     sortable: false,
//     align: "left",
//     filterable: true,
//   },
//   {
//     value: "updated_at",
//     label: "最終更新日",
//     sortable: false,
//     align: "left",
//     filterable: true,
//   },
// ];

import { CheckBox } from "@mui/icons-material";
import {
  Button,
  Checkbox,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableFooter,
  TableHead,
  TablePagination,
  TableRow,
  TableSortLabel,
} from "@mui/material";

interface Props<T> {
  rows: T[];
  columns: Column[];
}

export type Column = {
  value: string;
  label: string;
  sortable: boolean;
  align: "left" | "right" | "center";
  filterable: boolean;
};

export function DataTable<T>(props: Props<T>) {
  return (
    <>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell padding="checkbox">
                <Checkbox
                  color="primary"
                  indeterminate={false}
                  checked={false}
                  onChange={() => {}}
                  inputProps={{
                    "aria-label": "select all desserts",
                  }}
                />
              </TableCell>
              {props.columns.map((column) => {
                return (
                  <TableCell key={column.value}>
                    <TableSortLabel>{column.label}</TableSortLabel>
                  </TableCell>
                );
              })}
              <TableCell padding="checkbox"></TableCell>
              <TableCell padding="checkbox"></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {props.rows.map((row: T, index) => (
              <TableRow key={index}>
                <TableCell padding="checkbox">
                  <Checkbox
                    color="primary"
                    checked={false}
                    inputProps={{
                      "aria-labelledby": "test",
                    }}
                  />
                </TableCell>
                {props.columns.map((column, index) => {
                  const key = column.value as keyof typeof row;
                  return (
                    <TableCell key={column.value + index}>{row[key]}</TableCell>
                  );
                })}
                <TableCell padding="checkbox">
                  <Button onClick={() => {}}>編集</Button>
                </TableCell>
                <TableCell padding="checkbox">
                  <Button onClick={() => {}}>削除</Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
          <TableFooter></TableFooter>
        </Table>
      </TableContainer>
      {/* <TablePagination
        count={0}
        onPageChange={() => {}}
        page={1}
        rowsPerPage={1}
      /> */}
    </>
  );
}
