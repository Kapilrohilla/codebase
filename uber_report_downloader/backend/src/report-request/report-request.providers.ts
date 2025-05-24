import { Connection } from 'mongoose';
import { ReportRequestSchema } from '../schema/report-request.schema';

export const reportRequestProviders = [
  {
    provide: 'REPORT_REQUEST_MODEL',
    useFactory: (connection: Connection) =>
      connection.model('ReportRequest', ReportRequestSchema),
    inject: ['DATABASE_CONNECTION'],
  },
];
