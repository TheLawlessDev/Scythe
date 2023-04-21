package handlers

import (
	"Scythe/core/providers"
	"Scythe/core/utility"
	"context"
)

func (h *Handler) VideoHandler(ctx context.Context) error {
	videoProviders := providers.ListProvidersByCategory("video")
	chosenProvider := utility.VideoProviderPrompt(videoProviders)

	links, err := providers.Start(ctx, chosenProvider)
	if err != nil {
		return err
	}

	h.URLs = append(h.URLs, links...)

	return nil
}
