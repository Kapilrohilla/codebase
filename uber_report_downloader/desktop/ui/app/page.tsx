"use client"

import { useState, useEffect } from "react"
import {
  FluentProvider,
  webLightTheme,
  Tab,
  TabList,
  type TabValue,
  Button,
  Input,
  Dropdown,
  Option,
  Card,
  CardHeader,
  CardPreview,
  Text,
  Badge,
  Spinner,
  Body1,
  Caption1,
  makeStyles,
  tokens,
} from "@fluentui/react-components"
import {
  CalendarLtr24Regular,
  DocumentArrowDown24Regular,
  Clock24Regular,
  CheckmarkCircle24Regular,
  ErrorCircle24Regular,
} from "@fluentui/react-icons"

const useStyles = makeStyles({
  container: {
    padding: tokens.spacingVerticalL,
    maxWidth: "1200px",
    margin: "0 auto",
  },
  formSection: {
    display: "flex",
    flexDirection: "column",
    gap: tokens.spacingVerticalM,
    marginBottom: tokens.spacingVerticalXL,
  },
  formRow: {
    display: "flex",
    gap: tokens.spacingHorizontalM,
    alignItems: "end",
    flexWrap: "wrap",
  },
  formField: {
    display: "flex",
    flexDirection: "column",
    gap: tokens.spacingVerticalXS,
    minWidth: "200px",
  },
  historyList: {
    display: "flex",
    flexDirection: "column",
    gap: tokens.spacingVerticalS,
  },
  historyCard: {
    padding: tokens.spacingVerticalM,
  },
  cardContent: {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
  },
  statusSection: {
    display: "flex",
    alignItems: "center",
    gap: tokens.spacingHorizontalS,
  },
  actionSection: {
    display: "flex",
    gap: tokens.spacingHorizontalS,
    alignItems: "center",
  },
})

interface ReportRequest {
  id: string
  startDate: Date
  endDate: Date
  organizations: string[]
  status: "pending" | "generating" | "completed" | "failed"
  createdAt: Date
  completedAt?: Date
  downloadUrl?: string
}

const organizations = [
  { id: "org1", name: "Organization A" },
  { id: "org2", name: "Organization B" },
  { id: "org3", name: "Organization C" },
  { id: "org4", name: "Organization D" },
  { id: "org5", name: "Organization E" },
]

export default function Component() {
  const styles = useStyles()
  const [selectedTab, setSelectedTab] = useState<TabValue>("generate")
  const [startDate, setStartDate] = useState<Date | null>(null)
  const [endDate, setEndDate] = useState<Date | null>(null)
  const [selectedOrgs, setSelectedOrgs] = useState<string[]>([])
  const [requests, setRequests] = useState<ReportRequest[]>([])
  const [isGenerating, setIsGenerating] = useState(false)

  // Mock API functions
  const generateReport = async (startDate: Date, endDate: Date, organizations: string[]) => {
    // Simulate API call
    const newRequest: ReportRequest = {
      id: `req_${Date.now()}`,
      startDate,
      endDate,
      organizations,
      status: "pending",
      createdAt: new Date(),
    }

    setRequests((prev) => [newRequest, ...prev])
    return newRequest.id
  }

  const checkReportStatus = async (requestId: string): Promise<ReportRequest["status"]> => {
    // Simulate API call with random progression
    const request = requests.find((r) => r.id === requestId)
    if (!request) return "failed"

    if (request.status === "completed" || request.status === "failed") {
      return request.status
    }

    // Simulate status progression
    const random = Math.random()
    if (request.status === "pending" && random > 0.7) {
      return "generating"
    } else if (request.status === "generating" && random > 0.6) {
      return "completed"
    } else if (random < 0.05) {
      return "failed"
    }

    return request.status
  }

  const downloadReport = async (requestId: string) => {
    // Simulate download
    const request = requests.find((r) => r.id === requestId)
    if (request) {
      // In a real app, this would trigger an actual download
      alert(
        `Downloading report for ${request.organizations.join(", ")} from ${request.startDate.toLocaleDateString()} to ${request.endDate.toLocaleDateString()}`,
      )
    }
  }

  // Long polling for active requests
  useEffect(() => {
    const activeRequests = requests.filter((r) => r.status === "pending" || r.status === "generating")

    if (activeRequests.length === 0) return

    const pollInterval = setInterval(async () => {
      for (const request of activeRequests) {
        const newStatus = await checkReportStatus(request.id)

        setRequests((prev) =>
          prev.map((r) => {
            if (r.id === request.id && r.status !== newStatus) {
              return {
                ...r,
                status: newStatus,
                completedAt: newStatus === "completed" || newStatus === "failed" ? new Date() : r.completedAt,
                downloadUrl: newStatus === "completed" ? `https://api.example.com/download/${r.id}` : r.downloadUrl,
              }
            }
            return r
          }),
        )
      }
    }, 10000) // Poll every 10 seconds

    return () => clearInterval(pollInterval)
  }, [requests])

  const handleGenerate = async () => {
    if (!startDate || !endDate || selectedOrgs.length === 0) {
      alert("Please select date range and at least one organization")
      return
    }

    setIsGenerating(true)
    try {
      await generateReport(startDate, endDate, selectedOrgs)
      // Reset form
      setStartDate(null)
      setEndDate(null)
      setSelectedOrgs([])
    } catch (error) {
      alert("Failed to generate report")
    } finally {
      setIsGenerating(false)
    }
  }

  const getStatusIcon = (status: ReportRequest["status"]) => {
    switch (status) {
      case "pending":
        return <Clock24Regular />
      case "generating":
        return <Spinner size="tiny" />
      case "completed":
        return <CheckmarkCircle24Regular />
      case "failed":
        return <ErrorCircle24Regular />
    }
  }

  const getStatusBadge = (status: ReportRequest["status"]) => {
    const appearance = {
      pending: "outline" as const,
      generating: "tint" as const,
      completed: "filled" as const,
      failed: "outline" as const,
    }

    const color = {
      pending: "warning" as const,
      generating: "brand" as const,
      completed: "success" as const,
      failed: "danger" as const,
    }

    return (
      <Badge appearance={appearance[status]} color={color[status]}>
        {status.charAt(0).toUpperCase() + status.slice(1)}
      </Badge>
    )
  }

  return (
    <FluentProvider theme={webLightTheme}>
      <div className={styles.container}>
        <TabList selectedValue={selectedTab} onTabSelect={(_, data) => setSelectedTab(data.value)}>
          <Tab id="generate" value="generate">
            <CalendarLtr24Regular />
            Generate Report
          </Tab>
          <Tab id="history" value="history">
            <DocumentArrowDown24Regular />
            Request History
          </Tab>
        </TabList>

        {selectedTab === "generate" && (
          <div className={styles.formSection}>
            <Text as="h2" size={600}>
              Generate New Report
            </Text>

            <div className={styles.formRow}>
              <div className={styles.formField}>
                <Text as="label" weight="semibold">
                  Start Date
                </Text>
                <Input
                  type="date"
                  value={startDate ? startDate.toISOString().split("T")[0] : ""}
                  onChange={(_, data) => {
                    if (data.value) {
                      setStartDate(new Date(data.value))
                    } else {
                      setStartDate(null)
                    }
                  }}
                  placeholder="Select start date"
                />
              </div>

              <div className={styles.formField}>
                <Text as="label" weight="semibold">
                  End Date
                </Text>
                <Input
                  type="date"
                  value={endDate ? endDate.toISOString().split("T")[0] : ""}
                  onChange={(_, data) => {
                    if (data.value) {
                      setEndDate(new Date(data.value))
                    } else {
                      setEndDate(null)
                    }
                  }}
                  placeholder="Select end date"
                />
              </div>

              <div className={styles.formField}>
                <Text as="label" weight="semibold">
                  Organizations
                </Text>
                <Dropdown
                  multiselect
                  placeholder="Select organizations"
                  value={selectedOrgs.map((id) => organizations.find((org) => org.id === id)?.name || id).join(", ")}
                  selectedOptions={selectedOrgs}
                  onOptionSelect={(_, data) => {
                    if (data.optionValue) {
                      setSelectedOrgs((prev) =>
                        prev.includes(data.optionValue!)
                          ? prev.filter((id) => id !== data.optionValue)
                          : [...prev, data.optionValue!],
                      )
                    }
                  }}
                >
                  {organizations.map((org) => (
                    <Option key={org.id} value={org.id}>
                      {org.name}
                    </Option>
                  ))}
                </Dropdown>
              </div>

              <Button
                appearance="primary"
                disabled={isGenerating || !startDate || !endDate || selectedOrgs.length === 0}
                onClick={handleGenerate}
              >
                {isGenerating ? <Spinner size="tiny" /> : null}
                Generate Report
              </Button>
            </div>
          </div>
        )}

        {selectedTab === "history" && (
          <div className={styles.formSection}>
            <Text as="h2" size={600}>
              Request History
            </Text>

            {requests.length === 0 ? (
              <Card className={styles.historyCard}>
                <CardPreview>
                  <Text>No report requests found. Generate your first report to see it here.</Text>
                </CardPreview>
              </Card>
            ) : (
              <div className={styles.historyList}>
                {requests.map((request) => (
                  <Card key={request.id} className={styles.historyCard}>
                    <CardHeader
                      header={
                        <div className={styles.cardContent}>
                          <div>
                            <Body1>
                              {request.startDate.toLocaleDateString()} - {request.endDate.toLocaleDateString()}
                            </Body1>
                            <Caption1>
                              Organizations:{" "}
                              {request.organizations
                                .map((id) => organizations.find((org) => org.id === id)?.name || id)
                                .join(", ")}
                            </Caption1>
                            <Caption1>
                              Created: {request.createdAt.toLocaleString()}
                              {request.completedAt && ` • Completed: ${request.completedAt.toLocaleString()}`}
                            </Caption1>
                          </div>

                          <div className={styles.actionSection}>
                            <div className={styles.statusSection}>
                              {getStatusIcon(request.status)}
                              {getStatusBadge(request.status)}
                            </div>

                            {request.status === "completed" && (
                              <Button
                                appearance="primary"
                                icon={<DocumentArrowDown24Regular />}
                                onClick={() => downloadReport(request.id)}
                              >
                                Download
                              </Button>
                            )}
                          </div>
                        </div>
                      }
                    />
                  </Card>
                ))}
              </div>
            )}
          </div>
        )}
      </div>
    </FluentProvider>
  )
}
