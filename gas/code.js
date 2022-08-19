// function doGet(e) {
//   console.log(e);

//   return ContentService.createTextOutput(
//     JSON.stringify({status: 'ok'})    
//   ).setMimeType(ContentService.MimeType.JSON);
// }

function doPost(e) {
  console.log(e);

  const ss    = SpreadsheetApp.getActiveSpreadsheet();
  const sheet = ss.getSheetByName('Sheet1'); 
  const data  = JSON.parse(e.postData.contents);
  const result = appendDataToSpreadSheet(data, sheet);

  return ContentService
    .createTextOutput(JSON.stringify(result))
    .setMimeType(ContentService.MimeType.JSON);
}

function appendDataToSpreadSheet(data, sheet) {
  let result = {};

  if (!data.word) {
    result["error"] = "body json data doesn't contain required parameters.";
    result["status"] = "400";
    return result;
  };

  if (!data.created_at) {
    const now = new Date();
    data["created_at"] = Utilities.formatDate(now, 'JST', 'yyyy/MM/dd HH:mm:ss')
  }

  try {
    // 行の最後に値を追加
    sheet.appendRow([
      data.created_at,
      data.word,
      data.meaning || "",
      data.example || ""
    ]);
    result["status"] = "200";
  }
  catch (err) {
    result["error"] = err.message;
    result["status"] = "500";
  }
  
  return result;
}
