export class CreateReportDto {
  from: Date;
  to: Date;
  reportType: 'vehicle_performance' | 'driver_performance';
  organizations: string[];
}
