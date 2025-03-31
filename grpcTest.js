const { exec } = require('child_process');
const path = require('path');

const protoPath = path.join(__dirname, 'grpc/proto', 'user.proto');
const dataFilePath = path.join(__dirname, 'payload-small.json');
const command = `ghz --insecure --proto ${protoPath} --call userpb.PaymentService.ProcessUser --data-file ${dataFilePath} -c 100 --duration 20s localhost:50051`;

exec(command, (error, stdout, stderr) => {
  if (error) {
    console.error('❌ Error running ghz:', error.message);
    return;
  }
  if (stderr) {
    console.error('⚠️ stderr:', stderr);
  }

  console.log(stdout);
});
