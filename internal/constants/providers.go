package constants

import "github.com/qdm12/ddns-updater/internal/models"

// All possible provider values.
const (
	AZURE        models.Provider = "azure"
	CLOUDFLARE   models.Provider = "cloudflare"
	DIGITALOCEAN models.Provider = "digitalocean"
	DDNSSDE      models.Provider = "ddnss"
	DONDOMINIO   models.Provider = "dondominio"
	DNSOMATIC    models.Provider = "dnsomatic"
	DNSPOD       models.Provider = "dnspod"
	DUCKDNS      models.Provider = "duckdns"
	DYN          models.Provider = "dyn"
	DREAMHOST    models.Provider = "dreamhost"
	GODADDY      models.Provider = "godaddy"
	GOOGLE       models.Provider = "google"
	HE           models.Provider = "he"
	INFOMANIAK   models.Provider = "infomaniak"
	NAMECHEAP    models.Provider = "namecheap"
	NOIP         models.Provider = "noip"
	SELFHOSTDE   models.Provider = "selfhost.de"
)

func ProviderChoices() []models.Provider {
	return []models.Provider{
		AZURE,
		CLOUDFLARE,
		DIGITALOCEAN,
		DDNSSDE,
		DONDOMINIO,
		DNSOMATIC,
		DNSPOD,
		DUCKDNS,
		DYN,
		DREAMHOST,
		GODADDY,
		GOOGLE,
		HE,
		INFOMANIAK,
		NAMECHEAP,
		NOIP,
		SELFHOSTDE,
	}
}
