package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-transfer-family/mcp-server/config"
	"github.com/aws-transfer-family/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func StartfiletransferHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.StartFileTransferRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/#X-Amz-Target=TransferService.StartFileTransfer", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Target"]; ok {
			req.Header.Set("X-Amz-Target", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.StartFileTransferResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateStartfiletransferTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=TransferService_StartFileTransfer",
		mcp.WithDescription("<p>Begins a file transfer between local Amazon Web Services storage and a remote AS2 or SFTP server.</p> <ul> <li> <p>For an AS2 connector, you specify the <code>ConnectorId</code> and one or more <code>SendFilePaths</code> to identify the files you want to transfer.</p> </li> <li> <p>For an SFTP connector, the file transfer can be either outbound or inbound. In both cases, you specify the <code>ConnectorId</code>. Depending on the direction of the transfer, you also specify the following items:</p> <ul> <li> <p>If you are transferring file from a partner's SFTP server to a Transfer Family server, you specify one or more <code>RetreiveFilePaths</code> to identify the files you want to transfer, and a <code>LocalDirectoryPath</code> to specify the destination folder.</p> </li> <li> <p>If you are transferring file to a partner's SFTP server from Amazon Web Services storage, you specify one or more <code>SendFilePaths</code> to identify the files you want to transfer, and a <code>RemoteDirectoryPath</code> to specify the destination folder.</p> </li> </ul> </li> </ul>"),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("RetrieveFilePaths", mcp.Description("")),
		mcp.WithString("SendFilePaths", mcp.Description("")),
		mcp.WithString("ConnectorId", mcp.Required(), mcp.Description("")),
		mcp.WithString("LocalDirectoryPath", mcp.Description("")),
		mcp.WithString("RemoteDirectoryPath", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StartfiletransferHandler(cfg),
	}
}
