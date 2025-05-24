import { Schema, Document } from 'mongoose';

// Define the schema
export const ReportRequestSchema = new Schema(
  {
    fromDate: {
      type: Date,
      required: true,
    },
    toDate: {
      type: Date,
      required: true,
    },
    reportType: {
      type: String,
      required: true,
    },
    organizations: {
      type: [String],
      required: true,
    },
    status: {
      type: String,
      default: 'pending',
    },
    downloadUrl: {
      type: String,
      required: false,
    },
  },
  { timestamps: true },
);

// Define the interface for the document
export interface ReportRequest {
  fromDate: Date;
  toDate: Date;
  reportType: string;
  organizations: string[];
  status: string;
  downloadUrl?: string;
  createdAt?: Date;
  updatedAt?: Date;
}

// Export the mongoose model
export type ReportRequestDocument = ReportRequest & Document;
