import { Button } from "@mui/material";
import React from "react";

export const DeleteSection = () => {
  const deleteSection = () => {};

  return (
    <div>
      <Button variant="outlined" onClick={deleteSection}>
        選択削除
      </Button>
    </div>
  );
};
